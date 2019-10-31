package server

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
)

//RoomManager is a struct that keeps track of all the picto rooms.
type RoomManager struct {
	Rooms     map[string]*Room `json:"Rooms"`
	MaxRooms  int              `json:"MaxRooms"`
	RoomCount int              `json:"RoomCount"`
	apiToken  string
	Mode      string
}

//NewRoomManager creates a new room manager.
func NewRoomManager(MaxRooms int, apiToken string, Mode string) RoomManager {
	r := RoomManager{
		Rooms:     make(map[string]*Room, MaxRooms),
		MaxRooms:  MaxRooms,
		RoomCount: 0,
		apiToken:  apiToken,
		Mode:      Mode,
	}
	return r
}

func (rm *RoomManager) createRoom() (*Room, error) {
	if rm.RoomCount < rm.MaxRooms {
		newRoomID := strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
		for _, hasKey := rm.Rooms[newRoomID]; hasKey || newRoomID == ""; {
			newRoomID = strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
		}

		newRoom := newRoom(rm, newRoomID, MaxRoomSize)
		rm.Rooms[newRoom.ID] = newRoom
		rm.RoomCount++

		return newRoom, nil
	}
	return nil, errors.New("Reached MaxRooms Limit.")
}

func (rm *RoomManager) destroyRoom(roomID string) {
	log.Println("Destroying room ID" + roomID)

	room := rm.Rooms[roomID]

	delete(rm.Rooms, roomID)
	rm.RoomCount--

	room.destroy()
}
