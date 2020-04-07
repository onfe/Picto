package server

import (
	"time"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager       *RoomManager
	ID            string         `json:"ID"`
	Name          string         `json:"Name"`
	Static        bool           `json:"Static"`
	Public        bool           `json:"Public"`
	ClientManager *ClientManager `json:"ClientManager"`
	EventCache    *CircularQueue `json:"EventCache"`
	LastUpdate    time.Time      `json:"LastUpdate"`
	Closing       bool           `json:"Closing"`
	CloseTime     time.Time      `json:"CloseTime"`
}

func newRoom(manager *RoomManager, roomID string, name string, static bool, public bool, maxClients int) *Room {
	r := Room{
		manager:       manager,
		ID:            roomID,
		Name:          name,
		Static:        static,
		Public:        public,
		ClientManager: newClientManager(maxClients),
		EventCache:    newCircularQueue(ChatHistoryLen),
		LastUpdate:    time.Now(),
		Closing:       false,
	}
	return &r
}

func (r *Room) getDetails() string {
	return "(Room ID \"" + r.ID + "\" ('" + r.Name + "'))"
}

func (r *Room) getClientNames() []string {
	return r.ClientManager.getClientNames()
}

func (r *Room) addClient(c *Client) error {
	err := r.ClientManager.addClient(c)
	if err != nil {
		return err
	}

	//Now that the client has successfully been added to the server, the LastUpdate can be updated to now.
	r.LastUpdate = time.Now()

	/*
		2 * the apropriate min message interval is subtracted from the client's lastmessage time to ensure they
		can immediately send a message upon join.
	*/
	if r.Public {
		c.LastMessage = c.LastMessage.Add(-2 * MinMessageIntervalPublic)
	} else {
		c.LastMessage = c.LastMessage.Add(-2 * MinMessageInterval)
	}

	/*
		The client is sent an initialisation event, then all other clients are informed of the user's having joined the room.
		To do this, an array of strings of all the clients' usernames (including the new client's) has to be constructed.
	*/
	clientNames := r.getClientNames()
	clientNames[c.ID] = c.Name

	//Updating the new client as to the room state with an init event.
	c.sendBuffer <- newInitEvent(r.ID, r.Name, r.Static, c.ID, clientNames).toBytes()

	//Updating the new client with all the messages from the message cache.
	for _, E := range r.EventCache.getAll() {
		if E != nil {
			e := E.(*EventWrapper)
			//currentTime is UNIX time in millisecond precision.
			currentTime := time.Now().UnixNano() / int64(time.Millisecond)
			if !r.Public ||
				(r.Public &&
					(e.Time > currentTime-StaticMessageTimeout)) {
				c.sendBuffer <- e.toBytes()
			}
		}
	}

	//Now the new client is up to date and in the clients map of the room, all the clients are notified of their presence.
	r.distributeEvent(newUserEvent(c.ID, c.Name, clientNames), true, -1)

	return nil
}

func (r *Room) removeClient(clientID int) error {
	client := r.ClientManager.Clients[clientID]

	err := r.ClientManager.removeClient(clientID)
	if err != nil {
		return err
	}

	r.LastUpdate = time.Now()

	r.distributeEvent(newUserEvent(clientID, client.Name, r.getClientNames()), true, -1)

	return nil
}

func (r *Room) pruneClients() {
	r.ClientManager.pruneClients(ClientMessageTimeout)
}

func (r *Room) distributeEvent(event *EventWrapper, cached bool, sender int) {
	r.ClientManager.distributeEvent(event, sender)

	r.LastUpdate = time.Now()

	if cached {
		r.EventCache.push(event)
	}
}

func (r *Room) announce(message string) {
	r.distributeEvent(newAnnouncementEvent(message), true, -1)
}

func (r *Room) closeable() bool {
	switch true {
	case r.Static:
		return false
	case r.Closing:
		return time.Now().After(r.CloseTime)
	default:
		return r.ClientManager.ClientCount == 0 && time.Since(r.LastUpdate) > RoomGracePeriod
	}
}

func (r *Room) close() {
	r.ClientManager.closeClients()
}
