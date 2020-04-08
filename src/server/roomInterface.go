package server

import "time"

//RoomInterface must be implemented by all kinds of rooms.
type roomInterface interface {
	getID() string
	getType() string

	addClient(c *client) error
	removeClient(clientID int) error
	pruneClients()

	//recieveEvent handles events recieved from clients.
	recieveEvent(event *eventWrapper, sender *client)

	//announce announces a string message to all clients of the room.
	announce(message string)

	closeable() bool
	setCloseTime(time.Time)
	close()
}
