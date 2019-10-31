package server

import (
	"math/rand"
	"strconv"
)

//RoomManager is a struct that keeps track of all the picto rooms.
type RoomManager struct {
	Rooms     map[string]*Room `json:"Rooms"`
	MaxRooms  int              `json:"MaxRooms"`
	RoomCount int              `json:"RoomCount"`
}

func NewRoomManager(MaxRooms int) RoomManager {
	return RoomManager{
		Rooms:     make(map[string]*Room, MaxRooms),
		MaxRooms:  MaxRooms,
		RoomCount: 0,
	}
}

func (rm *RoomManager) createRoom() *Room {
	newRoomID := strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
	for _, hasKey := rm.Rooms[newRoomID]; hasKey || newRoomID == ""; {
		newRoomID = strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
	}
	newRoom := newRoom(rm, newRoomID, MaxRoomSize)
	rm.Rooms[newRoom.ID] = newRoom
	rm.RoomCount++
	return newRoom
}

func (rm *RoomManager) destroyRoom(roomID string) {
	room := rm.Rooms[roomID]
	delete(rm.Rooms, roomID)
	rm.RoomCount--
	room.destroy()
}
