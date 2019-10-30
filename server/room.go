package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager      *RoomManager
	id           string
	name         string
	clients      map[int]*Client
	maxClients   int
	messageCache CircularQueue
	lastUpdate   time.Time
}

func newRoom(manager *RoomManager, roomID string, maxClients int) Room {
	r := Room{
		id:           roomID,
		name:         "Picto Room",
		clients:      make(map[int]*Client),
		maxClients:   maxClients,
		messageCache: newCircularQueue(ChatHistoryLen),
		lastUpdate:   time.Now(),
	}

	return r
}

func (r *Room) addClient(c *Client) bool {
	if len(r.clients) < r.maxClients {
		for _, client := range r.clients {
			if client.name == c.name {
				return false
			}
		}
		r.lastUpdate = time.Now()
		for _, M := range r.messageCache.getAll() {
			if M != nil {
				m := M.(Message)
				c.send(websocket.TextMessage, m.body)
			}
		}
		r.clients[len(r.clients)] = c
		return true
	}
	return false
}

func (r *Room) removeClient(clientID int) {
	log.Println("Room:", r.id, "removed client:", clientID, ":", r.clients[clientID].name)
	r.lastUpdate = time.Now()
	client := r.clients[clientID]
	delete(r.clients, clientID)
	client.destroy()
	if len(r.clients) == 0 {
		r.manager.destroyRoom(r.id)
	}
}

func (r *Room) distributeMessage(m Message) {
	r.lastUpdate = time.Now()
	r.messageCache.push(m)
	for _, client := range r.clients {
		if client.id != m.senderID {
			client.sendBuffer <- m.body
		}
	}
}

func (r *Room) destroy() {
	for _, client := range r.clients {
		client.destroy()
	}
}
