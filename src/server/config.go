package server

import "time"

const (
	//MaxRooms = maximum amount of rooms the server may have at any one time
	MaxRooms = 10
	//DefaultRoomSize = Default size of default room.
	DefaultRoomSize = 8
	//MaxMessageSize = Max size of a message from the client.
	MaxMessageSize = 50000 //Max size of a picto image in Bytes
	//MinMessageInterval = Minimum interval between messages sent by a client to be acknowledged.
	MinMessageInterval = time.Second
	//ChatHistoryLen = Number of messages kept by server per room.
	ChatHistoryLen = 10
	//ClientSendTimeout is the time allotted for a message to be sent.
	ClientSendTimeout = 10 * time.Second
	//ClientTimeout = Max interval allotted between pings and pongs.
	ClientTimeout = 60 * time.Second
	//ClientPingPeriod is the period between pings.
	ClientPingPeriod = time.Second
)
