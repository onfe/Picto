package main

import (
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
	id          int
	parentRoom  *Room
	name        string
	ws          *websocket.Conn
	sendBuffer  chan []byte
	lastMessage time.Time
	lastPong    time.Time
}

func newClient(w http.ResponseWriter, r *http.Request, parentRoom *Room, name string) (Client, error) {
	ws, err := upgrader.Upgrade(w, r, nil)

	c := Client{
		id:          0,
		parentRoom:  parentRoom,
		name:        name,
		ws:          ws,
		sendBuffer:  make(chan []byte, 256),
		lastMessage: time.Now(),
		lastPong:    time.Now(),
	}

	if err == nil {
		go c.sendLoop()
		go c.recieveLoop()
	}

	return c, err
}

func (c *Client) sendLoop() {
	ticker := time.NewTicker(ClientPingPeriod)
	defer func() {
		ticker.Stop()
		c.destroy()
	}()
	for {
		var err error
		select {
		case message, ok := <-c.sendBuffer:
			if !ok {
				c.destroy()
			}
			err = c.send(websocket.TextMessage, message)
		case <-ticker.C:
			err = c.send(websocket.PingMessage, nil)
		}
		if err != nil {
			c.destroy()
		}
	}
}

func (c *Client) send(messageType int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(ClientSendTimeout))
	return c.ws.WriteMessage(messageType, payload)
}

func (c *Client) recieveLoop() {
	defer func() {
		c.destroy()
	}()

	c.ws.SetReadLimit(MaxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(ClientTimeout))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(ClientTimeout))
		return nil
	})

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.recieve(newMessage(message, c.id))
	}
}

func (c *Client) recieve(m Message) {
	if time.Since(c.lastMessage) > MinMessageInterval {
		c.parentRoom.distributeMessage(m)
	}
}

func (c *Client) destroy() {
	c.send(websocket.CloseMessage, []byte{})
	return
}
