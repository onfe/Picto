package server

import (
	"encoding/json"
	"log"
	"time"
)

//EventWrapper is the wrapper around every event
type EventWrapper struct {
	Event   string
	Time    int64
	Payload interface{}
}

func newEventWrapper(event string, payload interface{}) []byte {
	eventWrapper := EventWrapper{
		Event:   event,
		Time:    time.Now().Unix(),
		Payload: payload,
	}
	data, err := json.Marshal(eventWrapper)
	if err != nil {
		log.Println("[EVENTS] - Couldn't marshal new EventWrapper")
	}
	return data
}

func (e EventWrapper) getEventData() []byte {
	data, _ := json.Marshal(e)
	return data
}

//InitEvent is sent to clients when they join a room to inform them of the room's state.
type InitEvent struct {
	RoomID    string
	RoomName  string
	UserIndex int      //Index of the user that just joined in the users array.
	Users     []string //Array of strings of users' names.
}

func newInitEvent(roomID string, roomName string, userIndex int, users []string) []byte {
	initEvent := InitEvent{
		RoomID:    roomID,
		RoomName:  roomName,
		UserIndex: userIndex,
		Users:     users,
	}
	return newEventWrapper("init", initEvent)
}

//UserEvent is sent to clients to inform them of when another client leaves/joins their room.
type UserEvent struct {
	UserIndex int //Index of the user that just joined in the users array.
	Users     []string
}

func newUserEvent(userIndex int, users []string) []byte {
	userEvent := UserEvent{
		UserIndex: userIndex,
		Users:     users,
	}
	return newEventWrapper("user", userEvent)
}

//MessageEvent is sent to clients to inform them of a new message in their room.
type MessageEvent struct {
	ColourIndex int    //Index of the user that just sent the message
	Sender      string //Name of the user that sent the message (at the time it was recieved).
	Message     map[string]interface{}
}

//AnnouncementEvent is sent to clients to inform them of an announcement in the room.
type AnnouncementEvent struct {
	Announcement string
}

func newAnnouncementEvent(announcement string) []byte {
	announcementEvent := AnnouncementEvent{
		Announcement: announcement,
	}
	return newEventWrapper("announcement", announcementEvent)
}

//RenameEvent is sent to clients to inform them of the name of their room being changed.
type RenameEvent struct {
	UserIndex int
	RoomName  string
}
