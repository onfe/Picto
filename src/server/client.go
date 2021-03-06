package server

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, //REMOVE FOR PROD
}

//Client is a struct that contains all of the info about a client.
type client struct {
	room        roomInterface
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	IP          string
	ws          *websocket.Conn
	sendBuffer  chan []byte
	LastMessage time.Time `json:"LastMessage"`
	LastPong    time.Time `json:"LastPong"`
	closed      bool
}

func newClient(w http.ResponseWriter, r *http.Request, Name string, IP string) (*client, error) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return nil, err
	}

	c := client{
		Name:        Name,
		ws:          ws,
		IP:          IP,
		sendBuffer:  make(chan []byte, 256),
		LastMessage: time.Now(),
		LastPong:    time.Now(),
	}

	//'If a message read from the client exceeds this limit, the connection sends a close message to the peer and returns ErrReadLimit to the application'
	c.ws.SetReadLimit(MaxMessageSize)

	//If no message is recieved before this deadline, the websocket is corrupt and all future reads will return an error.
	c.ws.SetReadDeadline(time.Now().Add(ClientTimeout))

	//When a pong is recieved, the read deadline is updated.
	c.ws.SetPongHandler(func(string) error {
		c.LastPong = time.Now()
		c.ws.SetReadDeadline(time.Now().Add(ClientTimeout))
		return nil
	})

	c.ws.SetCloseHandler(func(code int, reason string) error {
		log.Println("[CLIENT] - Closed connection to "+c.getDetails()+" with code", code, "and reason:", reason)
		if c.room != nil {
			c.room.removeClient(c.ID)
		}
		c.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, reason))
		return c.ws.Close()
	})

	//Check the client was actually given a valid name.
	if len(Name) > MaxClientNameLength {
		err = errors.New("username too long")
	} else if Name == "" {
		err = errors.New("username not provided")
	}

	return &c, err
}

func (c *client) getDetails() string {
	if c.room != nil {
		return "(Room ID " + c.room.getID() + " : Client ID " + strconv.Itoa(c.ID) + ")"
	}
	return "(Roomless : Client ID " + strconv.Itoa(c.ID) + ")"
}

//Cancel should only be called before GO.
func (c *client) Cancel(code int, text string) {
	c.ws.CloseHandler()(code, text)
}

//GO starts the client's send and recieve goroutines.
func (c *client) GO() {
	go c.sendLoop()
	go c.recieveLoop()
}

func (c *client) close() {
	c.closed = true
}

//----------------------------------------------------------------------------------------------------SENDLOOP

func (c *client) sendLoop() {
	ticker := *time.NewTicker(ClientPingPeriod)

	//When the send loop loses connection to the client or sendBuffer is closed.
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case message, open := <-c.sendBuffer:
			if !open {
				log.Println("[CLIENT] - sendBuffer closed of " + c.getDetails())
				c.ws.CloseHandler()(websocket.CloseNormalClosure, "Connection Closed By Server")
				return
			}

			err := c.send(websocket.TextMessage, message)
			if err != nil {
				log.Println("[CLIENT] - Failed to distribute message to "+c.getDetails()+", error:", err.Error())
				c.ws.CloseHandler()(websocket.CloseGoingAway, "Internal Server Error #1")
				return
			}

			h := sha1.New()
			h.Write(message)
			log.Println("[CLIENT] - Distributed message to "+c.getDetails()+", byte string:", hex.EncodeToString(h.Sum(nil)))

		case <-ticker.C:
			if c.closed {
				close(c.sendBuffer)
			}
			err := c.send(websocket.PingMessage, nil)
			if err != nil {
				log.Println("[CLIENT] - Failed to send ping to "+c.getDetails()+", error:", err.Error())
				c.ws.CloseHandler()(websocket.CloseGoingAway, "Failed to ping")
				return
			}
		}
	}
}

func (c *client) send(messageType int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(ClientSendTimeout))
	return c.ws.WriteMessage(messageType, payload)
}

//----------------------------------------------------------------------------------------------------RECIEVELOOP

func (c *client) recieveLoop() {
	//Loops, pulling messages from the websocket.
	for {
		_, data, err := c.ws.ReadMessage()
		if err != nil {
			log.Println("[CLIENT] - Readloop got error from websocket connection and stopped:", err)
			c.closed = true
			break
		}
		event := &eventWrapper{}
		err = json.Unmarshal(data, event)
		if err != nil {
			log.Println("[CLIENT] - Readloop got an invalid message from " + c.getDetails() + ": " + string(data))
		} else {
			c.room.recieveEvent(event, c)
		}
	}
}

func (c *client) recieve(e *eventWrapper) {
	//Rate limiting: the client recieves no indication that their message was ignored due to rate limiting.
	if time.Since(c.LastMessage) > MinMessageInterval {
		c.LastMessage = time.Now()
		h := sha1.New()
		h.Write(e.toBytes())
		log.Println("[CLIENT] - Recieved message from "+c.getDetails()+", byte string:", hex.EncodeToString(h.Sum(nil)))
		c.room.recieveEvent(e, c)
	}
}
