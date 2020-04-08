package server

import (
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

type submission struct {
	Sender  string
	Message *MessageEvent
}

//SubmissionRoom is a struct that holds all the info about a singular picto SubmissionRoom.
type SubmissionRoom struct {
	manager *RoomManager

	ID          string `json:"ID"`
	Description string `json:"Description"`

	ClientManager *ClientManager `json:"ClientManager"`

	EventCache *CircularQueue `json:"EventCache"`

	SubmissionCache *CircularQueue `json:"Submissions"`
	Submittees      map[string]*submission

	LastUpdate time.Time `json:"LastUpdate"`
	Closing    bool      `json:"Closing"`
	CloseTime  time.Time `json:"CloseTime"`
}

func newSubmissionRoom(manager *RoomManager, name, description string, maxClients int) *SubmissionRoom {
	r := SubmissionRoom{
		manager:         manager,
		ID:              name,
		Description:     description,
		ClientManager:   newClientManager(maxClients),
		EventCache:      newCircularQueue(ChatHistoryLen),
		SubmissionCache: newCircularQueue(MaxSubmissions),
		Submittees:      make(map[string]*submission),
		LastUpdate:      time.Now(),
		Closing:         false,
	}
	return &r
}

//------------------------------ Utils ------------------------------
//distributeEvent is a handy wrapper to make event caching easier.
func (r *SubmissionRoom) distributeEvent(event *EventWrapper, cached bool, sender int) {
	r.ClientManager.distributeEvent(event, sender)

	r.LastUpdate = time.Now()

	if cached {
		r.EventCache.push(event)
	}
}

//------------------------------ Implementing RoomInterface ------------------------------

//The significant differences between rooms should lie in how they handle client events (in recieveEvents).
func (r *SubmissionRoom) recieveEvent(event *EventWrapper, sender *Client) {
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

		//...otherwise we fill in the ColourIndex and Sender fields.
		message.ColourIndex = sender.ID
		message.Sender = sender.Name

		// We then need to create a submission...
		addr := sender.ws.RemoteAddr().String()
		_, month, day := time.Now().Date()

		sub := &submission{
			Sender:  addr + "-" + strconv.Itoa(day) + "-" + month.String(),
			Message: &message,
		}

		//We add it to the submissioncache if they haven't already submitted.
		if _, alreadySubmitted := r.Submittees[sub.Sender]; !alreadySubmitted {
			sender.sendBuffer <- newAnnouncementEvent("Thank you for your submission!").toBytes()
			r.SubmissionCache.push(sub)
		} else {
			sender.sendBuffer <- newAnnouncementEvent("Thank you for your new submission! Your previous one has been overwritten.").toBytes()
		}
		/* and we always update the submission by either creating/updating it by
		reference through the Submittees map...
		*/
		r.Submittees[sub.Sender] = sub

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
		r.ID = rename.RoomName
		rename.UserName = sender.Name
		r.distributeEvent(wrapEvent("rename", rename), true, -1)
	}
}

func (r *SubmissionRoom) getID() string {
	return r.ID
}

func (r *SubmissionRoom) getType() string {
	return "submission"
}

func (r *SubmissionRoom) addClient(c *Client) error {
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

	//Updating the new client with all the messages from the message cache.
	for _, E := range r.EventCache.getAll() {
		if E != nil {
			e := E.(*EventWrapper)
			c.sendBuffer <- e.toBytes()
		}
	}

	c.sendBuffer <- newAnnouncementEvent(r.Description).toBytes()

	return nil
}

func (r *SubmissionRoom) removeClient(clientID int) error {
	return r.ClientManager.removeClient(clientID)
}

func (r *SubmissionRoom) pruneClients() {
	r.ClientManager.pruneClients(ClientMessageTimeout)
}

func (r *SubmissionRoom) announce(message string) {
	r.distributeEvent(newAnnouncementEvent(message), true, -1)
}

func (r *SubmissionRoom) closeable() bool {
	switch true {
	case r.Closing:
		return time.Now().After(r.CloseTime)
	default:
		return false
	}
}

func (r *SubmissionRoom) setCloseTime(closeTime time.Time) {
	r.CloseTime = closeTime
	r.Closing = true
}

func (r *SubmissionRoom) close() {
	r.ClientManager.closeClients()
}
