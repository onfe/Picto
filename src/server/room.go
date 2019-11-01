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

func (r *Room) getDetails() string {
	return "(Room ID" + r.ID + " ('" + r.Name + "'))"
}

func (r *Room) addClient(c *Client) error {
	if r.ClientCount < r.MaxClients {
		//ClientCount is immediately incremented so there's little chance of two people joining the room within a short time peroid causing the room to become overpopulated.
		r.ClientCount++

		//Checking that the client's desired name is not already taken.
		for _, client := range r.Clients {
			if client.Name == c.Name {
				//If it is, then ClientCount can be decremented as they've failed to join the room.
				r.ClientCount--
				return errors.New("Name already taken.")
			}
		}

		//The client is sent all of the messages currently in the MessageCache of the room.
		for _, M := range r.MessageCache.getAll() {
			if M != nil {
				m := M.(Message)
				c.send(websocket.TextMessage, m.Body)
			}
		}

		//Generating an ID for the new client.
		newClientID := strconv.Itoa(rand.Intn(r.MaxClients * (10 ^ 4)))
		for _, hasKey := r.Clients[newClientID]; hasKey || newClientID == ""; {
			newClientID = strconv.Itoa(rand.Intn(r.MaxClients * (10 ^ 4)))
		}

		r.Clients[newClientID] = c
		r.Clients[newClientID].ID = newClientID

		//Now that the client has successfully been added to the server, the LastUpdate can be updated to now.
		r.LastUpdate = time.Now()

		return nil
	}
	return errors.New("Room already full.")
}

func (r *Room) removeClient(clientID string) error {
	if client, exists := r.Clients[clientID]; exists {
		delete(r.Clients, clientID)
		r.ClientCount--

		r.LastUpdate = time.Now()
		log.Println("Removed client:", client.getDetails())

		if len(r.Clients) == 0 {
			log.Println("Closed empty room:", r.getDetails())
			r.manager.closeRoom(r.ID)
		}

		return nil
	}
	return errors.New("Room does not have such a client")
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

func (r *Room) close() {
	for _, client := range r.Clients {
		client.closeConnection("Room closed by server.")
	}
}
