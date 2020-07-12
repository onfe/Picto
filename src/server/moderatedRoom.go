package server

import (
	"errors"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

//ModeratedRoom is a struct that holds all the info about a singular picto ModeratedRoom.
type moderatedRoom struct {
	manager *RoomManager

	ID          string `json:"ID"`
	Description string `json:"Description"`

	ClientManager *clientManager `json:"ClientManager"`

	ModerationCache *moderationCache `json:"ModerationCache"`
	IgnoredIPs      map[string]time.Time

	Closing   bool      `json:"Closing"`
	CloseTime time.Time `json:"CloseTime"`
}

func newModeratedRoom(manager *RoomManager, name, description string, maxClients int) *moderatedRoom {
	r := moderatedRoom{
		manager:         manager,
		ID:              name,
		Description:     description,
		ClientManager:   newClientManager(maxClients),
		ModerationCache: newModerationCache(MaxMessages),
		IgnoredIPs:      make(map[string]time.Time),
		Closing:         false,
	}
	return &r
}

//------------------------------ Utils ------------------------------

func (r *moderatedRoom) setMessageState(messageID string, newState string) error {
	//update its state in the messages cache
	newID, err := r.ModerationCache.setState(messageID, newState) //should never return an error.
	if err != nil {
		return err
	}

	//If it's being made visible...
	if newState == visible {
		//...get the message from the cache...
		moderatedMessage, moderatedMessageExists := r.ModerationCache.Messages[newID] //~should~ never return an error
		if !moderatedMessageExists {
			return errors.New("could not find moderatedMessage with id: " + newID)
		}

		//...update its Time field to now...
		moderatedMessage.Message.Time = time.Now().UnixNano() / int64(time.Millisecond)

		//...and distribute it.
		r.ClientManager.distributeEvent(moderatedMessage.Message, -1)
	}

	return nil
}

func (r *moderatedRoom) deleteMessage(messageID string, offensive bool) error {
	moderatedMessage, err := r.ModerationCache.remove(messageID)
	if err != nil {
		return err
	}

	if offensive {
		ipSansPort := strings.Split(moderatedMessage.SenderIP, ":")[0]
		r.IgnoredIPs[ipSansPort] = time.Now()
	}

	return nil
}

//------------------------------ Implementing RoomInterface ------------------------------

//The significant differences between rooms should lie in how they handle client events (in recieveEvents).
func (r *moderatedRoom) recieveEvent(event *eventWrapper, sender *client) {
	switch event.Event {
	case "message":
		//First check if the client has been ignored...
		ipSansPort := strings.Split(sender.IP, ":")[0]
		ignoreTime, ignored := r.IgnoredIPs[ipSansPort]
		//If the client is ignored, check when they were ignored.
		if ignored {
			//If it was less than the ClientIgnoreTime, ignore the message...
			if time.Since(ignoreTime) < ClientIgnoreTime {
				return
			}
			//...otherwise, remove the IgnoredClients entry and continue.
			delete(r.IgnoredIPs, ipSansPort)
		}

		//The payload field of EventWrapper is defined as interface{},
		// Unmarshal throws the payload into a map[string]interface{}.
		// We need to decode it.
		message := messageEvent{}
		mapstructure.Decode(event.Payload, &message)

		//If the message is empty, we ignore it...
		if message.isEmpty() {
			return
		}

		//...otherwise we fill in the ColourIndex and Sender fields.
		message.ColourIndex = sender.ID
		message.Sender = sender.Name

		// We then need to wrap it and create a moderatedMessage...
		msg := &moderatedMessage{
			SenderIP: sender.IP,
			Message:  wrapEvent("message", message),
		}

		// ...add it to the moderation cache, and distribute.
		r.ModerationCache.add(msg)
		r.ClientManager.distributeEvent(msg.Message, sender.ID)
	}
}

func (r *moderatedRoom) getID() string {
	return r.ID
}

func (r *moderatedRoom) getType() string {
	return "moderated"
}

func (r *moderatedRoom) addClient(c *client) error {
	err := r.ClientManager.addClient(c)
	if err != nil {
		return err
	}

	/*
		2 * the apropriate min message interval is subtracted from the client's lastmessage time to ensure they
		can immediately send a message upon join.
	*/
	c.LastMessage = c.LastMessage.Add(-2 * MinMessageInterval)

	//Creating a fake users list with only the joining user in it...
	clientNames := make([]string, r.ClientManager.MaxClients)
	clientNames[c.ID] = c.Name

	//Updating the new client as to the room state with an init event.
	c.sendBuffer <- newInitEvent(r.ID, r.ID, true, c.ID, clientNames).toBytes()

	//Updating the new client with all the visible messages from the message cache.
	for _, s := range r.ModerationCache.getAll() {
		if s != nil {
			if s.State == visible {
				c.sendBuffer <- s.Message.toBytes()
			}
		}
	}

	//Provide the client with the room description as an announcement.
	c.sendBuffer <- newAnnouncementEvent(r.Description).toBytes()

	//Now the new client is up to date and in the clients map of the room, all the clients are notified of their presence.
	r.ClientManager.distributeEvent(newUserEvent(c.ID, c.Name, r.ClientManager.getClientNames()), -1)

	return nil
}

func (r *moderatedRoom) removeClient(clientID int) error {
	client := r.ClientManager.Clients[clientID]

	err := r.ClientManager.removeClient(clientID)
	if err != nil {
		return err
	}

	r.ClientManager.distributeEvent(newUserEvent(clientID, client.Name, r.ClientManager.getClientNames()), -1)

	return nil
}

func (r *moderatedRoom) pruneClients() {
	r.ClientManager.pruneClients(ClientMessageTimeout)
}

func (r *moderatedRoom) announce(message string) {
	r.ClientManager.distributeEvent(newAnnouncementEvent(message), -1)
}

func (r *moderatedRoom) closeable() bool {
	switch true {
	case r.Closing:
		return time.Now().After(r.CloseTime)
	default:
		return false
	}
}

func (r *moderatedRoom) setCloseTime(closeTime time.Time) {
	r.CloseTime = closeTime
	r.Closing = true
}

func (r *moderatedRoom) close() {
	r.ClientManager.closeClients()
}
