package server

import (
	"errors"
	"log"
	"time"
)

//ClientManager is responsible for managing its room's clients.
type clientManager struct {
	Clients     []*client `json:"Clients"`
	ClientCount int       `json:"ClientCount"`
	MaxClients  int       `json:"MaxClients"`
}

func newClientManager(maxClients int) *clientManager {
	return &clientManager{
		Clients:     make([]*client, maxClients),
		ClientCount: 0,
		MaxClients:  maxClients,
	}
}

func (cm *clientManager) getClientNames() []string {
	names := make([]string, cm.MaxClients)
	for i, user := range cm.Clients {
		if user != nil {
			names[i] = user.Name
		}
	}
	return names
}

func (cm *clientManager) addClient(c *client) error {
	if cm.ClientCount < cm.MaxClients {
		//ClientCount is immediately incremented so there's little chance of two people joining the room within a short time peroid causing the room to become overpopulated.
		cm.ClientCount++

		//Checking that the client's desired name is not already taken.
		for _, client := range cm.Clients {
			if client != nil && client.Name == c.Name {
				//If it is, then ClientCount can be decremented as they've failed to join the room.
				cm.ClientCount--
				return errors.New("username already taken in this room")
			}
		}

		//Generating an ID for the new client.
		newClientID := 0
		for cm.Clients[newClientID] != nil {
			//Modulo is added just in case some fucky asynchronisation causes us to run over the end of the list.
			newClientID = (newClientID + 1) % cm.MaxClients
		}

		//The new client is added to the room's clients array.
		cm.Clients[newClientID] = c
		cm.Clients[newClientID].ID = newClientID

		return nil
	}
	return errors.New("this room is full")
}

func (cm *clientManager) removeClient(clientID int) error {
	if cm.Clients[clientID] != nil {
		client := cm.Clients[clientID]
		cm.Clients[clientID] = nil

		log.Println("[CLIENT MANAGER] - Removed client:", client.getDetails())

		cm.ClientCount--

		return nil
	}
	return errors.New("room does not have such a client")
}

func (cm *clientManager) pruneClients(timeout time.Duration) {
	for _, client := range cm.Clients {
		if client != nil {
			if time.Since(client.LastMessage) > timeout {
				client.close()
			}
		}
	}
}

func (cm *clientManager) getClientByRemoteAddr(remoteAddr string) (*client, error) {
	for _, client := range cm.Clients {
		if client != nil {
			if client.ws.RemoteAddr().String() == remoteAddr {
				return client, nil
			}
		}
	}
	return nil, errors.New("client couldn't be found with remoteaddr: " + remoteAddr)
}

func (cm *clientManager) distributeEvent(event *eventWrapper, sender int) {
	for _, client := range cm.Clients {
		if client != nil && client.ID != sender {
			client.sendBuffer <- event.toBytes()
		}
	}
}

func (cm *clientManager) closeClients() {
	for _, client := range cm.Clients {
		if client != nil {
			client.close()
		}
	}
}
