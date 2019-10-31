package server

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager      *RoomManager
	ID           string             `json:"ID"`
	Name         string             `json:"Name"`
	Clients      map[string]*Client `json:"Clients"`
	ClientCount  int                `json:"ClientCount"`
	MaxClients   int                `json:"MaxClients"`
	MessageCache *CircularQueue     `json:"MessageCache"`
	LastUpdate   time.Time          `json:"LastUpdate"`
}

func newRoom(manager *RoomManager, roomID string, maxClients int) *Room {
	r := Room{
		manager:      manager,
		ID:           roomID,
		Name:         "Picto Room",
		Clients:      make(map[string]*Client),
		ClientCount:  0,
		MaxClients:   maxClients,
		MessageCache: newCircularQueue(ChatHistoryLen),
		LastUpdate:   time.Now(),
	}
	return &r
}

func (r *Room) addClient(c *Client) error {
	if r.ClientCount < r.MaxClients {
		for _, client := range r.Clients {
			if client.Name == c.Name {
				return errors.New("Name already taken.")
			}
		}
		r.LastUpdate = time.Now()
		for _, M := range r.MessageCache.getAll() {
			if M != nil {
				m := M.(Message)
				c.send(websocket.TextMessage, m.Body)
			}
		}

		r.ClientCount++

		newClientID := strconv.Itoa(rand.Intn(r.MaxClients * (10 ^ 4)))
		for _, hasKey := r.Clients[newClientID]; hasKey || newClientID == ""; {
			newClientID = strconv.Itoa(rand.Intn(r.MaxClients * (10 ^ 4)))
		}

		r.Clients[newClientID] = c
		r.Clients[newClientID].ID = newClientID

		return nil
	}
	return errors.New("Room already full.")
}

func (r *Room) removeClient(clientID string) {
	if r.Clients[clientID] != nil {
		log.Println("Room ID"+r.ID, "('"+r.Name+"') removed client:", clientID, "('"+r.Clients[clientID].Name+"')")

		r.LastUpdate = time.Now()

		client := r.Clients[clientID]
		client.closeConnection()
		delete(r.Clients, clientID)
		r.ClientCount--

		if len(r.Clients) == 0 {
			r.manager.destroyRoom(r.ID)
		}
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
		client.closeConnection()
	}
}
