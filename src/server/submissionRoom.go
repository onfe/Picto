package server

import (
	"errors"
	"time"

	"github.com/mitchellh/mapstructure"
)

//SubmissionRoom is a struct that holds all the info about a singular picto SubmissionRoom.
type submissionRoom struct {
	manager *RoomManager

	ID          string `json:"ID"`
	Description string `json:"Description"`

	ClientManager *clientManager `json:"ClientManager"`

	EventCache *circularQueue `json:"EventCache"`

	SubmissionCache *submissionCache `json:"Submissions"`

	LastUpdate time.Time `json:"LastUpdate"`
	Closing    bool      `json:"Closing"`
	CloseTime  time.Time `json:"CloseTime"`
}

func newSubmissionRoom(manager *RoomManager, name, description string, maxClients int) *submissionRoom {
	r := submissionRoom{
		manager:         manager,
		ID:              name,
		Description:     description,
		ClientManager:   newClientManager(maxClients),
		EventCache:      newCircularQueue(ChatHistoryLen),
		SubmissionCache: newSubmissionCache(MaxSubmissions),
		LastUpdate:      time.Now(),
		Closing:         false,
	}
	return &r
}

//------------------------------ Utils ------------------------------
//distributeEvent is a handy wrapper to make event caching easier.
func (r *submissionRoom) distributeEvent(event *eventWrapper, cached bool, sender int) {
	r.ClientManager.distributeEvent(event, sender)

	r.LastUpdate = time.Now()

	if cached {
		r.EventCache.push(event)
	}
}

func (r *submissionRoom) publishSubmission(sender string) error {
	submission, submissionExists := r.SubmissionCache.Submissions[sender]
	if !submissionExists {
		return errors.New("could not find submission from sender: " + sender)
	}

	//Wrap it in an event and distribute it...
	event := wrapEvent("message", submission.Message)
	r.distributeEvent(event, true, -1)

	//...then remove it from the submissions cache
	err := r.SubmissionCache.remove(submission.ID) //should never return an error.
	if err != nil {
		return err
	}

	client, err := r.ClientManager.getClientByRemoteAddr(submission.Addr)
	if err != nil {
		//returns an err if client can't be found
		return nil
	}

	client.sendBuffer <- newAnnouncementEvent("Your submission just got published!").toBytes()
	client.sendBuffer <- newAnnouncementEvent("You can now make a new submission.").toBytes()

	return nil
}

func (r *submissionRoom) rejectSubmission(sender string) error {
	//Returns an error if a submission from the sender specified couldn't be found.
	return r.SubmissionCache.remove(sender)
}

//------------------------------ Implementing RoomInterface ------------------------------

//The significant differences between rooms should lie in how they handle client events (in recieveEvents).
func (r *submissionRoom) recieveEvent(event *eventWrapper, sender *client) {
	switch event.Event {
	case "message":
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

		// We then need to create a submission...
		sub := &submission{
			Addr:    sender.ws.RemoteAddr().String(),
			Message: &message,
		}

		// ...and add it to the submission cache
		alreadySubmitted := r.SubmissionCache.add(sub)

		//We give a different announcement depending upon if they have already made a submission or not.
		sender.sendBuffer <- newAnnouncementEvent("Thank you for your submission!").toBytes()
		if !alreadySubmitted {
			sender.sendBuffer <- newAnnouncementEvent("You can overwrite your submission by sending a new one.").toBytes()
		} else {
			sender.sendBuffer <- newAnnouncementEvent("Your previous submission has been overwritten.").toBytes()
		}

	case "rename":
		//The payload field of EventWrapper is defined as interface{},
		// Unmarshal throws the payload into a map[string]interface{}.
		// We need to decode it.
		rename := renameEvent{}
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

func (r *submissionRoom) getID() string {
	return r.ID
}

func (r *submissionRoom) getType() string {
	return "submission"
}

func (r *submissionRoom) addClient(c *client) error {
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
			e := E.(*eventWrapper)
			c.sendBuffer <- e.toBytes()
		}
	}

	c.sendBuffer <- newAnnouncementEvent(r.Description).toBytes()

	clientAddr := c.ws.RemoteAddr().String()
	_, submissionExists := r.SubmissionCache.Submissions[genSubmissionID(clientAddr)]

	if submissionExists {
		c.sendBuffer <- newAnnouncementEvent("You've already made a submission today, but you can overwrite it by sending another.").toBytes()
	}

	return nil
}

func (r *submissionRoom) removeClient(clientID int) error {
	return r.ClientManager.removeClient(clientID)
}

func (r *submissionRoom) pruneClients() {
	r.ClientManager.pruneClients(ClientMessageTimeout)
}

func (r *submissionRoom) announce(message string) {
	r.distributeEvent(newAnnouncementEvent(message), true, -1)
}

func (r *submissionRoom) closeable() bool {
	switch true {
	case r.Closing:
		return time.Now().After(r.CloseTime)
	default:
		return false
	}
}

func (r *submissionRoom) setCloseTime(closeTime time.Time) {
	r.CloseTime = closeTime
	r.Closing = true
}

func (r *submissionRoom) close() {
	r.ClientManager.closeClients()
}
