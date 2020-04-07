package server

import "time"

//RoomInterface must be implemented by all kinds of rooms.
type RoomInterface interface {
	getID() string
	getClientNames() []string
	getMinMessageInterval() time.Duration

	addClient(c *Client) error
	removeClient(clientID int) error
	pruneClients()

	renameable() bool
	rename(newName string)

	distributeEvent(event *EventWrapper, cached bool, sender int)
	announce(message string)

	closeable() bool
	setCloseTime(time.Time)
	close()
}
