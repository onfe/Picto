package server

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager      *RoomManager
	ID           string         `json:"ID"`
	Name         string         `json:"Name"`
	Clients      []*Client      `json:"Clients"`
	ClientCount  int            `json:"ClientCount"`
	MaxClients   int            `json:"MaxClients"`
	MessageCache *CircularQueue `json:"MessageCache"`
	LastUpdate   time.Time      `json:"LastUpdate"`
}

func newRoom(manager *RoomManager, roomID string, maxClients int) *Room {
	r := Room{
		manager:      manager,
		ID:           roomID,
		Name:         "Picto Room",
		Clients:      make([]*Client, maxClients),
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
			if client != nil && client.Name == c.Name {
				//If it is, then ClientCount can be decremented as they've failed to join the room.
				r.ClientCount--
				return errors.New("Name already taken.")
			}
		}

		//Generating an ID for the new client.
		newClientID := 0
		for r.Clients[newClientID] != nil {
			//Modulo is added just in case some fucky asynchronisation causes us to run over the end of the list.
			newClientID = (newClientID + 1) % r.MaxClients
		}

		//Now that the client has successfully been added to the server, the LastUpdate can be updated to now.
		r.LastUpdate = time.Now()

		/*
			The client is sent an initialisation event, then all other clients are informed of the user's having joined the room.
			To do this, an array of strings of all the clients' usernames (including the new client's) has to be constructed.
		*/
		clientNames := make([]string, r.MaxClients)
		clientNames[newClientID] = c.Name
		for i := 0; i < r.MaxClients; i++ {
			if r.Clients[i] != nil {
				clientNames[i] = r.Clients[i].Name
			}
		}

		//Updating the new client as to the room state with an init event.
		initEvent, _ := json.Marshal(
			InitEvent{
				Event:     Event{Event: "init"},
				RoomID:    r.ID,
				RoomName:  r.Name,
				UserIndex: newClientID,
				Users:     clientNames,
				NumUsers:  r.ClientCount,
			})
		c.sendBuffer <- initEvent

		//Updating the new client with all the messages from the message cache.
		for _, M := range r.MessageCache.getAll() {
			if M != nil {
				m := M.(Message)
				c.sendBuffer <- m.getEventData()
			}
		}

		//Now the new client is up to date, all the other clients are notified of their presence.
		for _, cc := range r.Clients {
			if cc != nil {
				userEvent, _ := json.Marshal(
					UserEvent{
						Event:     Event{Event: "user"},
						UserIndex: newClientID,
						Users:     clientNames,
						NumUsers:  r.ClientCount,
					})
				cc.sendBuffer <- userEvent
			}
		}

		//The new client is added to the room's clients array.
		r.Clients[newClientID] = c
		r.Clients[newClientID].ID = newClientID

		return nil
	}
	return errors.New("Room already full.")
}

func (r *Room) removeClient(clientID int) error {
	if r.Clients[clientID] != nil {
		client := r.Clients[clientID]
		r.Clients[clientID] = nil

		r.LastUpdate = time.Now()
		log.Println("Removed client:", client.getDetails())

		r.ClientCount--
		if r.ClientCount == 0 {
			r.ClientCount--
		}

		return nil
	}
	return errors.New("Room does not have such a client")
}

func (r *Room) distributeMessage(m Message) {
	r.LastUpdate = time.Now()
	r.MessageCache.push(m)
	messageData := m.getEventData()
	for _, client := range r.Clients {
		if client != nil && client.ID != m.SenderID {
			client.sendBuffer <- messageData
		}
	}
}

func (r *Room) closeAllConnections() {
	for _, client := range r.Clients {
		client.closeConnection("Room closed by server.")
	}
}
