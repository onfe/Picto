package server

import (
	"encoding/json"
	"log"
	"time"
)

//EventWrapper is the wrapper around every event
type eventWrapper struct {
	Event   string
	Time    int64
	Payload interface{}
}

func wrapEvent(event string, payload interface{}) *eventWrapper {
	eventWrapper := eventWrapper{
		Event:   event,
		Time:    time.Now().UnixNano() / int64(time.Millisecond),
		Payload: payload,
	}
	return &eventWrapper
}

func (e eventWrapper) toBytes() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		log.Println("[EVENTS] - Couldn't marshal new EventWrapper")
	}
	return data
}

//InitEvent is sent to clients when they join a room to inform them of the room's state.
type initEvent struct {
	RoomID    string
	RoomName  string
	Public    bool
	UserIndex int      //Index of the user that just joined in the users array.
	Users     []string //Array of strings of users' names.
}

func newInitEvent(roomID string, roomName string, public bool, userIndex int, users []string) *eventWrapper {
	initEvent := initEvent{
		RoomID:    roomID,
		RoomName:  roomName,
		Public:    public,
		UserIndex: userIndex,
		Users:     users,
	}
	return wrapEvent("init", initEvent)
}

//UserEvent is sent to clients to inform them of when another client leaves/joins their room.
//Including UserName might seem a bit redundant but it's neccessary when sending cached join/leave events.
type userEvent struct {
	UserIndex int    //Index of the user that just joined/left in the users array.
	UserName  string //Name of the user that just joined/left
	Users     []string
}

func newUserEvent(userIndex int, userName string, users []string) *eventWrapper {
	userEvent := userEvent{
		UserIndex: userIndex,
		UserName:  userName,
		Users:     users,
	}
	return wrapEvent("user", userEvent)
}

//MessageEvent is sent to clients to inform them of a new message in their room.
type messageEvent struct {
	ColourIndex int    //Index of the user that just sent the message
	Sender      string //Name of the user that sent the message (at the time it was recieved).
	Data        string //Message image, base64 encoded, RLE'd
	Span        int    //Width of the message
}

func (m *messageEvent) isEmpty() bool {
	return false //Dummy method, pretends all messages are not empty.
}

//AnnouncementEvent is sent to clients to inform them of an announcement in the room.
type announcementEvent struct {
	Announcement string
}

func newAnnouncementEvent(announcement string) *eventWrapper {
	announcementEvent := announcementEvent{
		Announcement: announcement,
	}
	return wrapEvent("announcement", announcementEvent)
}

//RenameEvent is sent to clients to inform them of the name of their room being changed.
type renameEvent struct {
	UserName string
	RoomName string
}

func newRenameEvent(userName string, roomName string) *eventWrapper {
	renameEvent := renameEvent{
		UserName: userName,
		RoomName: roomName,
	}
	return wrapEvent("rename", renameEvent)
}
