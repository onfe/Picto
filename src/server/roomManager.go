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
	Rooms           map[string]RoomInterface `json:"Rooms"`
	MaxRooms        int                      `json:"MaxRooms"`
	apiToken        string
	Mode            string
	wordsList       []string
	StaticRooms     map[string]*StaticRoom
	SubmissionRooms map[string]*SubmissionRoom
}

//NewRoomManager creates a new room manager.
func NewRoomManager(MaxRooms int, apiToken string, Mode string, wordsList []string, publicRoomConfigVar string) RoomManager {
	rm := RoomManager{
		Rooms:           make(map[string]RoomInterface, MaxRooms),
		MaxRooms:        MaxRooms,
		apiToken:        apiToken,
		Mode:            Mode,
		wordsList:       wordsList,
		StaticRooms:     make(map[string]*StaticRoom, MaxRooms),
		SubmissionRooms: make(map[string]*SubmissionRoom, MaxRooms),
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
			rm.createRoom(r.Name, r.Cap, true)
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
				if room.closeable() {
					rm.closeRoom(roomID)
				} else {
					room.pruneClients()
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

func (rm *RoomManager) createRoom(newRoomName string, maxClients int, static bool) (RoomInterface, error) {
	if len(rm.Rooms) < rm.MaxRooms {
		var room RoomInterface

		if !static {
			room = newRoom(rm, rm.generateNewRoomID(), newRoomName, maxClients)
		} else {
			/* We've got to check that static room names don't contain hyphens to
			ensure there is no overlap between dynamic and static rooms. */
			if strings.Contains(newRoomName, "-") {
				return nil, errors.New("static room names may not contain hyphens")
			}

			/* We've also got to check there's no static room already using the
			desired name anyway */
			for roomID := range rm.Rooms {
				if roomID == newRoomName {
					return nil, errors.New("a static room already exists with that name")
				}
			}

			room = newStaticRoom(rm, newRoomName, maxClients)
			rm.StaticRooms[room.getID()] = room.(*StaticRoom)
		}

		rm.Rooms[room.getID()] = room

		log.Println("[ROOM CREATED] - Created room ID \""+room.getID()+"\" | There are now", len(rm.Rooms), "active rooms.")

		return room, nil
	}
	return nil, errors.New("the server is at maximum rooms capacity")
}

func (rm *RoomManager) closeRoom(roomID string) {
	rm.Rooms[roomID].close()

	if _, isStatic := rm.StaticRooms[roomID]; isStatic {
		delete(rm.StaticRooms, roomID)
	}

	delete(rm.Rooms, roomID)

	log.Println("[ROOM CLOSED] - Closed room ID \""+roomID+"\" | There are now", len(rm.Rooms), "active rooms.")
}
