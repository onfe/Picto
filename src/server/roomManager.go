package server

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

//RoomManager is a struct that keeps track of all the picto rooms.
type RoomManager struct {
	Rooms     map[string]*Room `json:"Rooms"`
	MaxRooms  int              `json:"MaxRooms"`
	apiToken  string
	Mode      string
	wordsList []string
}

//NewRoomManager creates a new room manager.
func NewRoomManager(MaxRooms int, apiToken string, Mode string, wordsList []string) RoomManager {
	rm := RoomManager{
		Rooms:     make(map[string]*Room, MaxRooms),
		MaxRooms:  MaxRooms,
		apiToken:  apiToken,
		Mode:      Mode,
		wordsList: wordsList,
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
				if room.ClientCount == -1 && !room.Static {
					rm.closeRoom(roomID)
				}
			}

		}
	}
}

func (rm *RoomManager) generateNewRoomID() string {
	sep := "-"
	genID := func() string {
		return rm.wordsList[rand.Intn(len(rm.wordsList))] +
			sep + rm.wordsList[rand.Intn(len(rm.wordsList))] +
			sep + rm.wordsList[rand.Intn(len(rm.wordsList))]
	}
	newRoomID := genID()
	for _, taken := rm.Rooms[newRoomID]; taken; {
		newRoomID = genID()
	}
	return newRoomID
}

func (rm *RoomManager) createRoom(roomName string, static bool, maxClients int) (*Room, error) {
	if len(rm.Rooms) < rm.MaxRooms {
		newRoomID := rm.generateNewRoomID()
		newRoom := newRoom(rm, newRoomID, roomName, static, maxClients)
		rm.Rooms[newRoom.ID] = newRoom

		log.Println("[ROOM CREATED] - Created room ID \""+newRoomID+"\" | There are now", len(rm.Rooms), "active rooms.")
		return newRoom, nil
	}
	return nil, errors.New("Reached MaxRooms Limit.")
}

func (rm *RoomManager) closeRoom(roomID string) {
	if rm.Rooms[roomID].ClientCount > 0 {
		rm.Rooms[roomID].closeAllConnections()
	}
	delete(rm.Rooms, roomID)
	log.Println("[ROOM CLOSED] - Closed room ID \""+roomID+"\" | There are now", len(rm.Rooms), "active rooms.")
}
