package server

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//RoomManager is a struct that keeps track of all the picto rooms.
type RoomManager struct {
	Rooms       map[string]*Room `json:"Rooms"`
	MaxRooms    int              `json:"MaxRooms"`
	apiToken    string
	Mode        string
	wordsList   []string
	StaticRooms map[string]*Room
}

//NewRoomManager creates a new room manager.
func NewRoomManager(MaxRooms int, apiToken string, Mode string, wordsList []string, publicRoomConfigVar string) RoomManager {
	rm := RoomManager{
		Rooms:       make(map[string]*Room, MaxRooms),
		MaxRooms:    MaxRooms,
		apiToken:    apiToken,
		Mode:        Mode,
		wordsList:   wordsList,
		StaticRooms: make(map[string]*Room, MaxRooms),
	}
	rm.loadStaticRoomConfig(publicRoomConfigVar)
	go rm.roomMonitorLoop()
	return rm
}

func (rm *RoomManager) loadStaticRoomConfig(varname string) {
	config, configured := os.LookupEnv(varname)
	if configured {
		type roomConfig struct {
			Name   string
			Cap    int
			Public bool
		}

		var roomConfigs []roomConfig
		configBytes := []byte(config)

		err := json.Unmarshal(configBytes, &roomConfigs)
		if err != nil {
			log.Println("[SYSTEM] - Couldn't unmarshal room config:", err)
		}

		for _, r := range roomConfigs {
			rm.createRoom(r.Name, r.Cap, true, r.Public)
		}
	} else {
		log.Println("[SYSTEM] - Couldn't find public room config env var.")
	}
}

func (rm *RoomManager) roomMonitorLoop() {
	ticker := *time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			for roomID, room := range rm.Rooms {
				if (!room.Static &&
					(room.ClientCount == -1 && time.Since(room.LastUpdate) > RoomGracePeriod) ||
					time.Since(room.LastUpdate) > RoomTimeout) ||
					(room.Closing && time.Now().After(room.CloseTime)) {
					rm.closeRoom(roomID)
				} else {
					for _, client := range room.Clients {
						if client != nil {
							if time.Since(client.LastMessage) > ClientMessageTimeout {
								client.close()
							}
						}
					}
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

func (rm *RoomManager) createRoom(roomName string, maxClients int, static, public bool) (*Room, error) {
	if len(rm.Rooms) < rm.MaxRooms {
		newRoomID := roomName
		if !static {
			newRoomID = rm.generateNewRoomID()
		} else {
			if strings.Contains(newRoomID, "-") {
				return nil, errors.New("static room names may not contain hyphens")
			}
			for roomID := range rm.Rooms {
				if roomID == newRoomID {
					return nil, errors.New("a static room already exists with that name")
				}
			}
		}
		newRoom := newRoom(rm, newRoomID, roomName, static, public, maxClients)
		rm.Rooms[newRoom.ID] = newRoom
		rm.StaticRooms[newRoomID] = newRoom

		log.Println("[ROOM CREATED] - Created room ID \""+newRoomID+"\" | There are now", len(rm.Rooms), "active rooms.")
		return newRoom, nil
	}
	return nil, errors.New("the server is at maximum rooms capacity")
}

func (rm *RoomManager) closeRoom(roomID string) {
	if rm.Rooms[roomID].ClientCount > 0 {
		rm.Rooms[roomID].close()
	}
	if rm.Rooms[roomID].Static {
		delete(rm.StaticRooms, roomID)
	}
	delete(rm.Rooms, roomID)
	log.Println("[ROOM CLOSED] - Closed room ID \""+roomID+"\" | There are now", len(rm.Rooms), "active rooms.")
}
