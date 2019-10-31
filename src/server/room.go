package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager      *RoomManager `json:"-"`
	ID           string
	Name         string
	Clients      map[int]*Client
	MaxClients   int
	MessageCache *CircularQueue
	LastUpdate   time.Time
}

func newRoom(manager *RoomManager, roomID string, maxClients int) Room {
	r := Room{
		ID:           roomID,
		Name:         "Picto Room",
		Clients:      make(map[int]*Client),
		MaxClients:   maxClients,
		MessageCache: newCircularQueue(ChatHistoryLen),
		LastUpdate:   time.Now(),
	}

	return r
}

func (r *Room) addClient(c *Client) bool {
	if len(r.Clients) < r.MaxClients {
		for _, client := range r.Clients {
			if client.Name == c.Name {
				return false
			}
		}
		r.LastUpdate = time.Now()
		for _, M := range r.MessageCache.getAll() {
			if M != nil {
				m := M.(Message)
				c.send(websocket.TextMessage, m.Body)
			}
		}
		r.Clients[len(r.Clients)] = c
		return true
	}
	return false
}

func (r *Room) removeClient(clientID int) {
	log.Println("Room:", r.ID, "removed client:", clientID, ":", r.Clients[clientID].Name)
	r.LastUpdate = time.Now()
	client := r.Clients[clientID]
	delete(r.Clients, clientID)
	client.destroy()
	if len(r.Clients) == 0 {
		r.manager.destroyRoom(r.ID)
	}
}

func (r *Room) distributeMessage(m Message) {
	r.LastUpdate = time.Now()
	r.MessageCache.push(m)
	for _, client := range r.Clients {
		if client.ID != m.SenderID {
			client.sendBuffer <- m.Body
		}
	}
}

func (r *Room) destroy() {
	for _, client := range r.Clients {
		client.destroy()
	}
}
