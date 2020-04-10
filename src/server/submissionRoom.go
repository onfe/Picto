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

	SubmissionCache *submissionCache `json:"Submissions"`

	Closing   bool      `json:"Closing"`
	CloseTime time.Time `json:"CloseTime"`
}

func newSubmissionRoom(manager *RoomManager, name, description string, maxClients int) *submissionRoom {
	r := submissionRoom{
		manager:         manager,
		ID:              name,
		Description:     description,
		ClientManager:   newClientManager(maxClients),
		SubmissionCache: newSubmissionCache(MaxSubmissions),
		Closing:         false,
	}
	return &r
}

//------------------------------ Utils ------------------------------

func (r *submissionRoom) setSubmissionState(submissionID string, newState string) error {
	//update its state in the submissions cache
	newID, err := r.SubmissionCache.setState(submissionID, newState) //should never return an error.
	if err != nil {
		return err
	}

	//If it's being published...
	if newState == published {
		//...get the submission from the cache...
		submission, submissionExists := r.SubmissionCache.Submissions[newID] //~should~ never return an error
		if !submissionExists {
			return errors.New("could not find submission with id: " + newID)
		}

		//...update its Time field to now...
		submission.Message.Time = time.Now().UnixNano() / int64(time.Millisecond)

		//...distribute it...
		r.ClientManager.distributeEvent(submission.Message, -1)

		//...and, if they're still connected, congratulate the client.
		client, err := r.ClientManager.getClientByRemoteAddr(submission.Addr)
		if err != nil {
			//returns an err if client can't be found
			return nil
		}
		client.sendBuffer <- newAnnouncementEvent("Your submission just got published!").toBytes()
		client.sendBuffer <- newAnnouncementEvent("You can now make a new submission.").toBytes()
	}

	return nil
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

		// We then need to wrap it and create a submission...
		sub := &submission{
			Addr:    sender.ws.RemoteAddr().String(),
			Message: wrapEvent("message", message),
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
	for _, s := range r.SubmissionCache.getAll() {
		if s != nil {
			if s.State == published {
				c.sendBuffer <- s.Message.toBytes()
			}
		}
	}

	c.sendBuffer <- newAnnouncementEvent(r.Description).toBytes()

	clientAddr := c.ws.RemoteAddr().String()
	_, submissionExists := r.SubmissionCache.Submissions[r.SubmissionCache.genSubmissionID(clientAddr, submitted)]

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
	r.ClientManager.distributeEvent(newAnnouncementEvent(message), -1)
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
