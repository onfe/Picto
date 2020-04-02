package server

import "time"

const (
	//MaxRooms = maximum amount of rooms the server may have at any one time
	MaxRooms = 1024
	//RoomTimeout = maximum amount of time a room may stay open without any activity before being automatically closed
	RoomTimeout = 600 * time.Second
	//ClientMessageTimeout = maximum amount of time a client may be in a room without sending any messages before being disconnected
	ClientMessageTimeout = 300 * time.Second
	//StaticMessageTimeout = maximum amount of time a message exists in the cache of a static room before being ignored in milliseconds
	StaticMessageTimeout = 300 * 1000
	//DefaultRoomSize = Default size of default room.
	DefaultRoomSize = 8
	//MaxClientsPerRoom = max number of clients in a single room.
	MaxClientsPerRoom = 64
	//MaxRoomNameLength = the max length of a room name.
	MaxRoomNameLength = 32
	//MaxClientNameLength = the max length of a client name.
	MaxClientNameLength = 16
	//MaxMessageSize = Max size of a picto image in Bytes from the client.
	MaxMessageSize = 16384 //One byte per pixel in a 192*64 pixel canvas + a bit extra for the wrapper and other data.
	//MinMessageInterval = Minimum interval between messages sent by a client to be acknowledged.
	MinMessageInterval = time.Second
	//MinMessageIntervalStatic = Minimum interval between messages sent by a client to be acknowledged in a static room.
	MinMessageIntervalStatic = time.Second * 5
	//ChatHistoryLen = Number of messages kept by server per room.
	ChatHistoryLen = 64
	//ClientSendTimeout is the time allotted for a message to be sent.
	ClientSendTimeout = 10 * time.Second
	//ClientTimeout = Max interval allotted between pings and pongs.
	ClientTimeout = 60 * time.Second
	//ClientPingPeriod is the period between pings.
	ClientPingPeriod = time.Second
)
