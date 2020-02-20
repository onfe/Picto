package server

import (
	"errors"
	"log"
	"time"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager     *RoomManager
	ID          string         `json:"ID"`
	Name        string         `json:"Name"`
	Static      bool           `json:"Static"`
	Clients     []*Client      `json:"Clients"`
	ClientCount int            `json:"ClientCount"`
	MaxClients  int            `json:"MaxClients"`
	EventCache  *CircularQueue `json:"EventCache"`
	LastUpdate  time.Time      `json:"LastUpdate"`
}

func newRoom(manager *RoomManager, roomID string, name string, static bool, maxClients int) *Room {
	r := Room{
		manager:     manager,
		ID:          roomID,
		Name:        "Picto Room",
		Static:      static,
		Clients:     make([]*Client, maxClients),
		ClientCount: 0,
		MaxClients:  maxClients,
		EventCache:  newCircularQueue(ChatHistoryLen),
		LastUpdate:  time.Now(),
	}
	return &r
}

func (r *Room) getDetails() string {
	return "(Room ID \"" + r.ID + "\" ('" + r.Name + "'))"
}

func (r *Room) getClientNames() []string {
	names := make([]string, r.MaxClients)
	for i, user := range r.Clients {
		if user != nil {
			names[i] = user.Name
		}
	}
	return names
}

func (r *Room) changeName(event EventWrapper, nameChangerID int) {
	//Get the new room name and update it. This method might not be particularly safe?
	r.Name = (event.Payload.(map[string]interface{}))["RoomName"].(string)
	//Distribute the rename event to all the users.
	r.distributeEvent(event.toBytes(), false, nameChangerID)
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
				return errors.New("name already taken")
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
		clientNames := r.getClientNames()
		clientNames[newClientID] = c.Name

		//Updating the new client as to the room state with an init event.
		c.sendBuffer <- newInitEvent(r.ID, r.Name, newClientID, clientNames)

		//Updating the new client with all the messages from the message cache.
		for _, M := range r.EventCache.getAll() {
			if M != nil {
				c.sendBuffer <- M.([]byte)
			}
		}

		//Sending the client a "welcome to picto" announcement.
		c.sendBuffer <- newAnnouncementEvent("Welcome to Picto!")

		//The new client is added to the room's clients array.
		r.Clients[newClientID] = c
		r.Clients[newClientID].ID = newClientID

		//Now the new client is up to date and in the clients map of the room, all the clients are notified of their presence.
		r.distributeEvent(newUserEvent(newClientID, c.Name, clientNames), true, -1)

		return nil
	}
	return errors.New("room already full")
}

func (r *Room) removeClient(clientID int) error {
	if r.Clients[clientID] != nil {
		client := r.Clients[clientID]
		r.Clients[clientID] = nil

		r.LastUpdate = time.Now()
		log.Println("[ROOM] - Removed client:", client.getDetails())

		r.ClientCount--
		if r.ClientCount == 0 {
			r.ClientCount--
		}

		r.distributeEvent(newUserEvent(clientID, client.Name, r.getClientNames()), true, -1)
		return nil
	}
	return errors.New("Room does not have such a client")
}

func (r *Room) distributeEvent(event []byte, cached bool, sender int) {
	r.LastUpdate = time.Now()
	if cached {
		r.EventCache.push(event)
	}
	for _, client := range r.Clients {
		if client != nil && client.ID != sender {
			client.sendBuffer <- event
		}
	}
}

func (r *Room) announce(message string) {
	r.distributeEvent(newAnnouncementEvent(message), true, -1)
}

func (r *Room) closeAllConnections() {
	for _, client := range r.Clients {
		client.closeConnection("Room closed by server.")
	}
}
