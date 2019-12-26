package server

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
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
type Client struct {
	room        *Room
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	ws          *websocket.Conn
	sendBuffer  chan []byte
	LastMessage time.Time `json:"LastMessage"`
	LastPong    time.Time `json:"LastPong"`
	closeReason string
}

func newClient(w http.ResponseWriter, r *http.Request, Name string) (*Client, error) {
	ws, err := upgrader.Upgrade(w, r, nil)

	c := Client{
		Name:       Name,
		ws:         ws,
		sendBuffer: make(chan []byte, 256),
		//LastMessage is initially set 2x before the current time, to be sure the client can immediately start sending messages.
		LastMessage: time.Now().Add(-2 * MinMessageInterval),
		LastPong:    time.Now(),
	}

	if err == nil {
		go c.sendLoop()
		go c.recieveLoop()
	}

	return &c, err
}

func (c *Client) getDetails() string {
	if c.room != nil {
		return "(Room ID \"" + c.room.ID + "\" ('" + c.room.Name + "'): Client ID" + strconv.Itoa(c.ID) + " ('" + c.Name + "'))"
	}
	return "(Roomless: Client ID" + strconv.Itoa(c.ID) + " ('" + c.Name + "'))"
}

func (c *Client) closeConnection(reason string) {
	c.closeReason = reason
	close(c.sendBuffer)
}

//----------------------------------------------------------------------------------------------------SENDLOOP

func (c *Client) sendLoop() {
	ticker := *time.NewTicker(ClientPingPeriod)

	//When the send loop loses connection to the client or sendBuffer is closed.
	defer func() {
		ticker.Stop()

		log.Println("[CLIENT] - Sending close message to", c.Name, "for reason:", c.closeReason)
		c.send(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, c.closeReason))
		c.ws.Close()
	}()

	for {
		select {
		case message, open := <-c.sendBuffer:
			if !open {
				log.Println("[CLIENT] - Failed to get message from sendBuffer belonging to " + c.getDetails())
				return
			}

			err := c.send(websocket.TextMessage, message)
			if err != nil {
				log.Println("[CLIENT] - Failed to distribute message to "+c.getDetails()+", error:", err.Error())
				return
			}

			h := sha1.New()
			h.Write(message)
			log.Println("[CLIENT] - Distributed message to "+c.getDetails()+", byte string:", hex.EncodeToString(h.Sum(nil)))

		case <-ticker.C:
			err := c.send(websocket.PingMessage, nil)
			if err != nil {
				log.Println("[CLIENT] - Failed to send ping to "+c.getDetails()+", error:", err.Error())
				return
			}
		}
	}
}

func (c *Client) send(messageType int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(ClientSendTimeout))
	return c.ws.WriteMessage(messageType, payload)
}

//----------------------------------------------------------------------------------------------------RECIEVELOOP

func (c *Client) recieveLoop() {
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

	//When a close message is recieved, the client is removed from the room (if it's in one).
	c.ws.SetCloseHandler(func(code int, reason string) error {
		log.Println("[CLIENT] - Closed connection to "+c.getDetails()+" with code", code, "and reason:", reason)
		if c.closeReason == "" {
			c.closeReason = reason
			close(c.sendBuffer)
		}
		if c.room != nil {
			c.room.removeClient(c.ID)
		}
		return nil
	})

	//Loops, pulling messages from the websocket.
	for {
		_, data, err := c.ws.ReadMessage()
		if err != nil {
			log.Println("[CLIENT] - Readloop got error from websocket connection and stopped:", err)
			break
		}
		event := make(map[string]interface{})
		json.Unmarshal(data, &event)
		if _, valid := event["Event"]; !valid {
			log.Println("[CLIENT] - Readloop got an invalid message from " + c.getDetails() + ": " + string(data))
		} else {
			switch event["Event"] {
			case "message":
				var e MessageEvent
				json.Unmarshal(data, &e)
				e.UserIndex = c.ID
				c.recieve(e)
			case "rename":
				var e RenameEvent
				json.Unmarshal(data, &e)
				c.room.changeName(e.RoomName)
			}
		}
	}
}

func (c *Client) recieve(e Event) {
	//Rate limiting: the client recieves no indication that their message was ignored due to rate limiting.
	if time.Since(c.LastMessage) > MinMessageInterval {
		h := sha1.New()
		h.Write(e.getEventData())
		log.Println("[CLIENT] - Recieved message from "+c.getDetails()+", byte string:", hex.EncodeToString(h.Sum(nil)))
		c.room.distributeEvent(e)
	}
}
