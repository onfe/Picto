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

func wrapEvent(event string, payload interface{}) []byte {
	eventWrapper := EventWrapper{
		Event:   event,
		Time:    time.Now().UnixNano() / int64(time.Millisecond),
		Payload: payload,
	}
	return eventWrapper.toBytes()
}

func (e EventWrapper) toBytes() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		log.Println("[EVENTS] - Couldn't marshal new EventWrapper")
	}
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
	return wrapEvent("init", initEvent)
}

//UserEvent is sent to clients to inform them of when another client leaves/joins their room.
//Including UserName might seem a bit redundant but it's neccessary when sending cached join/leave events.
type UserEvent struct {
	UserIndex int    //Index of the user that just joined/left in the users array.
	UserName  string //Name of the user that just joined/left
	Users     []string
}

func newUserEvent(userIndex int, userName string, users []string) []byte {
	userEvent := UserEvent{
		UserIndex: userIndex,
		UserName:  userName,
		Users:     users,
	}
	return wrapEvent("user", userEvent)
}

//MessageEvent is sent to clients to inform them of a new message in their room.
type MessageEvent struct {
	ColourIndex int    //Index of the user that just sent the message
	Sender      string //Name of the user that sent the message (at the time it was recieved).
	Data        string //Message image, base64 encoded, RLE'd
	Span        int    //Width of the message
}

func (m *MessageEvent) isEmpty() bool {
	return false //Dummy method, pretends all messages are not empty.
}

//AnnouncementEvent is sent to clients to inform them of an announcement in the room.
type AnnouncementEvent struct {
	Announcement string
}

func newAnnouncementEvent(announcement string) []byte {
	announcementEvent := AnnouncementEvent{
		Announcement: announcement,
	}
	return wrapEvent("announcement", announcementEvent)
}

//RenameEvent is sent to clients to inform them of the name of their room being changed.
type RenameEvent struct {
	UserName string
	RoomName string
}

func newRenameEvent(userName string, roomName string) []byte {
	renameEvent := RenameEvent{
		UserName: userName,
		RoomName: roomName,
	}
	return wrapEvent("rename", renameEvent)
}
