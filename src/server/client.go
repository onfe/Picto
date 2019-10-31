package server

import (
	"log"
	"net/http"
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
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	ws          *websocket.Conn
	sendBuffer  chan []byte
	LastMessage time.Time `json:"LastMessage"`
	LastPong    time.Time `json:"LastPong"`
}

func newClient(w http.ResponseWriter, r *http.Request, room *Room, ID string, Name string) (*Client, error) {
	ws, err := upgrader.Upgrade(w, r, nil)

	c := Client{
		ID:          ID,
		room:        room,
		Name:        Name,
		ws:          ws,
		sendBuffer:  make(chan []byte, 256),
		LastMessage: time.Now(),
		LastPong:    time.Now(),
	}

	if err == nil {
		go c.sendLoop()
		go c.recieveLoop()
	}

	return &c, err
}

func (c *Client) sendLoop() {
	ticker := time.NewTicker(ClientPingPeriod)
	defer func() {
		ticker.Stop()
		log.Println("Send loop cut connection to client '" + c.Name + "' of room ID" + c.room.ID)
		c.room.removeClient(c.ID)
	}()

	for {
		var err error
		select {
		case message, ok := <-c.sendBuffer:
			if !ok {
				return
			}

			err = c.send(websocket.TextMessage, message)
			if err != nil {
				log.Println("Failed to distribute message to '"+c.Name+"' in room ID"+c.room.ID+", error:", err)
			} else {
				log.Println("Distributed message to '"+c.Name+"' in room ID"+c.room.ID+":", message)
			}

		case <-ticker.C:
			err = c.send(websocket.PingMessage, nil)
		}

		if err != nil {
			return
		}
	}
}

func (c *Client) send(messageType int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(ClientSendTimeout))
	return c.ws.WriteMessage(messageType, payload)
}

func (c *Client) recieveLoop() {
	defer func() {
		log.Println("Recieve loop cut connection to client '" + c.Name + "' of room ID" + c.room.ID)
		c.room.removeClient(c.ID)
	}()

	c.ws.SetReadLimit(MaxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(ClientTimeout))
	c.ws.SetPongHandler(func(string) error {
		c.LastPong = time.Now()
		c.ws.SetReadDeadline(time.Now().Add(ClientTimeout))
		return nil
	})

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.recieve(newMessage(message, c.ID))
	}
}

func (c *Client) recieve(m Message) {
	if time.Since(c.LastMessage) > MinMessageInterval {
		log.Println("Recieved message from '"+c.Name+"' (ID:"+c.ID+") in room ID"+c.room.ID+":", string(m.Body))
		c.room.distributeMessage(m)
	}
}

func (c *Client) closeConnection() {
	c.send(websocket.CloseMessage, []byte{})
}
