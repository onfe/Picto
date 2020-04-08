package server

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

//Room is a struct that holds all the info about a singular picto room.
type Room struct {
	manager *RoomManager

	ID   string `json:"ID"`
	Name string `json:"Name"`

	ClientManager *ClientManager `json:"ClientManager"`

	EventCache *CircularQueue `json:"EventCache"`

	LastUpdate time.Time `json:"LastUpdate"`
	Closing    bool      `json:"Closing"`
	CloseTime  time.Time `json:"CloseTime"`
}

func newRoom(manager *RoomManager, roomID string, name string, maxClients int) *Room {
	r := Room{
		manager: manager,

		ID:   roomID,
		Name: name,

		ClientManager: newClientManager(maxClients),
		EventCache:    newCircularQueue(ChatHistoryLen),

		LastUpdate: time.Now(),
		Closing:    false,
	}
	return &r
}

//------------------------------ Utils ------------------------------
//distributeEvent is a handy wrapper to make event caching easier.
func (r *Room) distributeEvent(event *EventWrapper, cached bool, sender int) {
	r.ClientManager.distributeEvent(event, sender)

	r.LastUpdate = time.Now()

	if cached {
		r.EventCache.push(event)
	}
}

//------------------------------ Implementing RoomInterface ------------------------------

//The significant differences between rooms should lie in how they handle client events (in recieveEvents).
func (r *Room) recieveEvent(event *EventWrapper, sender *Client) {
	switch event.Event {
	case "message":
		//The payload field of EventWrapper is defined as interface{},
		// Unmarshal throws the payload into a map[string]interface{}.
		// We need to decode it.
		message := MessageEvent{}
		mapstructure.Decode(event.Payload, &message)

		//If the message is empty, we ignore it...
		if message.isEmpty() {
			return
		}

		//...otherwise we fill in the ColourIndex and Sender fields,
		// rewrap it and recieve it.
		message.ColourIndex = sender.ID
		message.Sender = sender.Name
		r.distributeEvent(wrapEvent("message", message), true, sender.ID)

	case "rename":
		//The payload field of EventWrapper is defined as interface{},
		// Unmarshal throws the payload into a map[string]interface{}.
		// We need to decode it.
		rename := RenameEvent{}
		mapstructure.Decode(event.Payload, &rename)

		//If the new name is too long, we ignore it...
		if len(rename.RoomName) > MaxRoomNameLength {
			return
		}

		//...otherwise we change the room's name,
		// fill in the UserName field, rewrap it and distribute it...
		r.Name = rename.RoomName
		rename.UserName = sender.Name
		r.distributeEvent(wrapEvent("rename", rename), true, -1)
	}
}

func (r *Room) getID() string {
	return r.ID
}

func (r *Room) getType() string {
	return "dynamic"
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
	c.LastMessage = c.LastMessage.Add(-2 * MinMessageInterval)

	/*
		The client is sent an initialisation event, then all other clients are informed of the user's having joined the room.
		To do this, an array of strings of all the clients' usernames (including the new client's) has to be constructed.
	*/
	clientNames := r.ClientManager.getClientNames()
	clientNames[c.ID] = c.Name

	//Updating the new client as to the room state with an init event.
	c.sendBuffer <- newInitEvent(r.ID, r.Name, false, c.ID, clientNames).toBytes()

	//Updating the new client with all the messages from the message cache.
	for _, E := range r.EventCache.getAll() {
		if E != nil {
			e := E.(*EventWrapper)
			c.sendBuffer <- e.toBytes()
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

	r.distributeEvent(newUserEvent(clientID, client.Name, r.ClientManager.getClientNames()), true, -1)

	return nil
}

func (r *Room) pruneClients() {
	r.ClientManager.pruneClients(ClientMessageTimeout)
}

func (r *Room) announce(message string) {
	r.distributeEvent(newAnnouncementEvent(message), true, -1)
}

func (r *Room) closeable() bool {
	switch true {
	case r.Closing:
		return time.Now().After(r.CloseTime)
	default:
		return r.ClientManager.ClientCount == 0 && time.Since(r.LastUpdate) > RoomGracePeriod
	}
}

func (r *Room) setCloseTime(closeTime time.Time) {
	r.CloseTime = closeTime
	r.Closing = true
}

func (r *Room) close() {
	r.ClientManager.closeClients()
}
