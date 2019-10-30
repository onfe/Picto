package main

import (
	"math/rand"
	"strconv"
)

//RoomManager is a struct that keeps track of all the picto rooms.
type RoomManager struct {
	rooms     map[string]Room
	maxRooms  int
	roomCount int
}

func newRoomManager(maxRooms int) RoomManager {
	return RoomManager{
		rooms:     make(map[string]Room, MaxRooms),
		maxRooms:  maxRooms,
		roomCount: 0,
	}
}

func (rm RoomManager) createRoom() *Room {
	newRoomID := strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
	for _, hasKey := rm.rooms[newRoomID]; hasKey || newRoomID == ""; {
		newRoomID = strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
	}
	newRoom := newRoom(newRoomID, MaxRoomSize)
	rm.rooms[newRoom.id] = newRoom
	return &newRoom
}

func (rm RoomManager) destroyRoom(roomID string) {
	room := rm.rooms[roomID]
	delete(rm.rooms, roomID)
	rm.roomCount--
	room.destroy()
}
