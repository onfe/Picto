package server

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"
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
	rm := RoomManager{
		Rooms:     make(map[string]*Room, MaxRooms),
		MaxRooms:  MaxRooms,
		RoomCount: 0,
		apiToken:  apiToken,
		Mode:      Mode,
	}
	go rm.roomMonitorLoop()
	return rm
}

func (rm *RoomManager) roomMonitorLoop() {
	ticker := *time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			for roomID, room := range rm.Rooms {
				if room.ClientCount == -1 {
					log.Println("Closed empty room:", room.getDetails())
					rm.closeRoom(roomID)
					rm.RoomCount--
				}
			}

		}
	}
}

func (rm *RoomManager) createRoom() (*Room, error) {
	if rm.RoomCount < rm.MaxRooms {
		//RoomCount is immediately incremented so there's little chance of two people creating rooms within a short period of time causing there to become more than MaxRooms rooms.
		rm.RoomCount++

		newRoomID := strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
		for _, hasKey := rm.Rooms[newRoomID]; hasKey || newRoomID == ""; {
			newRoomID = strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
		}

		newRoom := newRoom(rm, newRoomID, MaxRoomSize)
		rm.Rooms[newRoom.ID] = newRoom

		return newRoom, nil
	}
	return nil, errors.New("Reached MaxRooms Limit.")
}

func (rm *RoomManager) closeRoom(roomID string) {
	log.Println("Closing room ID" + roomID)
	if rm.Rooms[roomID].ClientCount > 0 {
		rm.Rooms[roomID].closeAllConnections()
	}
	delete(rm.Rooms, roomID)
	rm.RoomCount--
}