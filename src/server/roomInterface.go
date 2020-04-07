package server

import "time"

//RoomInterface must be implemented by all kinds of rooms.
type RoomInterface interface {
	getID() string

	addClient(c *Client) error
	removeClient(clientID int) error
	pruneClients()

	//recieveEvent handles events recieved from clients.
	recieveEvent(event *EventWrapper, sender *Client)

	//announce announces a string message to all clients of the room.
	announce(message string)

	closeable() bool
	setCloseTime(time.Time)
	close()
}
