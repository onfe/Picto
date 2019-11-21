package server

//Event is the base struct for an event sent to clients.
type Event struct {
	event string
}

//InitEvent is sent to clients when they join a room to inform them of the room's state.
type InitEvent struct {
	Event
	RoomID    string
	RoomName  string
	UserIndex int      //Index of the user that just joined in the users array.
	Users     []string //Array of strings of users' names.
	NumUsers  int
}

//UserEvent is sent to clients to inform them of when another client leaves/joins their room.
type UserEvent struct {
	Event
	UserIndex int //Index of the user that just joined in the users array.
	Users     []string
	NumUsers  int
}

//MessageEvent is sent to clients to inform them of a new message in their room.
type MessageEvent struct {
	Event
	UserIndex int //Index of the user that just sent the message
	Message   []byte
}

//AnnouncementEvent is sent to clients to inform them of an announcement in the room.
type AnnouncementEvent struct {
	Event
	Announcement string
}

//RenameEvent is sent to clients to inform them of the name of their room being changed.
type RenameEvent struct {
	Event
	RoomName string
}
