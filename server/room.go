package main

import (
	"time"

	"github.com/gorilla/websocket"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager      *RoomManager
	id           string
	customName   string
	clients      map[int]Client
	clientCount  int
	maxClients   int
	messageCache CircularQueue
	lastUpdate   time.Time
}

func newRoom(roomID string, maxClients int) Room {
	r := Room{
		id:           roomID,
		customName:   "Picto Room",
		clients:      make(map[int]Client, MaxRoomSize),
		clientCount:  0,
		maxClients:   maxClients,
		messageCache: newCircularQueue(ChatHistoryLen),
		lastUpdate:   time.Now(),
	}
	return r
}

func (r *Room) addClient(c Client) {
	if r.clientCount < r.maxClients {
		r.clientCount++
		r.lastUpdate = time.Now()
		for _, M := range r.messageCache.getAll() {
			if M != nil {
				m := M.(Message)
				c.send(websocket.TextMessage, m.body)
			}
		}
		r.clients[c.id] = c
	}
}

func (r *Room) removeClient(clientID int) {
	r.lastUpdate = time.Now()
	client := r.clients[clientID]
	delete(r.clients, clientID)
	r.clientCount--
	client.destroy()
	if r.clientCount == 0 {
		r.manager.destroyRoom(r.id)
	}
}

func (r *Room) distributeMessage(m Message) {
	r.lastUpdate = time.Now()
	r.messageCache.push(m)
	for _, client := range r.clients {
		if client.id != m.senderID {
			client.send(websocket.TextMessage, m.body)
		}
	}
}

func (r *Room) destroy() {
	for _, client := range r.clients {
		client.destroy()
	}
}
