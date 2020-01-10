package server

import "encoding/json"

//Event is the base struct for an event sent to clients.
type Event interface {
	getEventData() []byte
	getEventType() string
	getSenderID() int
}

//InitEvent is sent to clients when they join a room to inform them of the room's state.
type InitEvent struct {
	Event     string
	RoomID    string
	RoomName  string
	UserIndex int      //Index of the user that just joined in the users array.
	Users     []string //Array of strings of users' names.
	NumUsers  int
}

func (e InitEvent) getEventData() []byte {
	data, _ := json.Marshal(e)
	return data
}
func (e InitEvent) getEventType() string { return e.Event }
func (e InitEvent) getSenderID() int     { return e.UserIndex }

//UserEvent is sent to clients to inform them of when another client leaves/joins their room.
type UserEvent struct {
	Event     string
	UserIndex int //Index of the user that just joined in the users array.
	Users     []string
	NumUsers  int
}

func (e UserEvent) getEventData() []byte {
	data, _ := json.Marshal(e)
	return data
}
func (e UserEvent) getEventType() string { return e.Event }
func (e UserEvent) getSenderID() int     { return e.UserIndex }

//MessageEvent is sent to clients to inform them of a new message in their room.
type MessageEvent struct {
	Event       string
	ColourIndex int    //Index of the user that just sent the message; used to set the message colour.
	Sender      string //Name of the user that sent the message (at the time it was recieved).
	Message     map[string]interface{}
}

func (e MessageEvent) getEventData() []byte {
	data, _ := json.Marshal(e)
	return data
}
func (e MessageEvent) getEventType() string { return e.Event }
func (e MessageEvent) getSenderID() int     { return e.ColourIndex }

//AnnouncementEvent is sent to clients to inform them of an announcement in the room.
type AnnouncementEvent struct {
	Event        string
	Announcement string
}

func (e AnnouncementEvent) getEventData() []byte {
	data, _ := json.Marshal(e)
	return data
}
func (e AnnouncementEvent) getEventType() string { return e.Event }
func (e AnnouncementEvent) getSenderID() int     { return -1 }

//RenameEvent is sent to clients to inform them of the name of their room being changed.
type RenameEvent struct {
	Event     string
	RoomName  string
	UserIndex int
}

func (e RenameEvent) getEventData() []byte {
	data, _ := json.Marshal(e)
	return data
}
func (e RenameEvent) getEventType() string { return e.Event }
func (e RenameEvent) getSenderID() int     { return e.UserIndex }
