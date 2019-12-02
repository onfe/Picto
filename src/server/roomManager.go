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
	Rooms    map[string]*Room `json:"Rooms"`
	MaxRooms int              `json:"MaxRooms"`
	apiToken string
	Mode     string
}

//NewRoomManager creates a new room manager.
func NewRoomManager(MaxRooms int, apiToken string, Mode string) RoomManager {
	rm := RoomManager{
		Rooms:    make(map[string]*Room, MaxRooms),
		MaxRooms: MaxRooms,
		apiToken: apiToken,
		Mode:     Mode,
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
					rm.closeRoom(roomID)
				}
			}

		}
	}
}

func (rm *RoomManager) createRoom() (*Room, error) {
	if len(rm.Rooms) < rm.MaxRooms {
		newRoomID := strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
		for _, hasKey := rm.Rooms[newRoomID]; hasKey || newRoomID == ""; {
			newRoomID = strconv.Itoa(rand.Intn(MaxRooms * (10 ^ 4)))
		}
		newRoom := newRoom(rm, newRoomID, MaxRoomSize)
		rm.Rooms[newRoom.ID] = newRoom

		log.Println("Created room ID"+newRoomID, "| There are now", len(rm.Rooms), "active rooms.")
		return newRoom, nil
	}
	return nil, errors.New("Reached MaxRooms Limit.")
}

func (rm *RoomManager) closeRoom(roomID string) {
	if rm.Rooms[roomID].ClientCount > 0 {
		rm.Rooms[roomID].closeAllConnections()
	}
	delete(rm.Rooms, roomID)
	log.Println("Closed room ID"+roomID, "| There are now", len(rm.Rooms), "active rooms.")
}
