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
	PublicRooms []room
}

type room struct {
	Name string
	Cap  int
}

//NewRoomManager creates a new room manager.
func NewRoomManager(MaxRooms int, apiToken string, Mode string, wordsList []string, publicRoomConfigVar string) RoomManager {
	rm := RoomManager{
		Rooms:     make(map[string]*Room, MaxRooms),
		MaxRooms:  MaxRooms,
		apiToken:  apiToken,
		Mode:      Mode,
		wordsList: wordsList,
	}
	rm.loadPublicRoomConfig(publicRoomConfigVar)
	go rm.roomMonitorLoop()
	return rm
}

func (rm *RoomManager) loadPublicRoomConfig(varname string) {
	config, configured := os.LookupEnv(varname)
	if configured {
		var rooms []room
		configBytes := []byte(config)

		err := json.Unmarshal(configBytes, &rooms)
		if err != nil {
			log.Println("[SYSTEM] - Couldn't unmarshal room config:", err)
		}

		for _, r := range rooms {
			rm.createRoom(r.Name, r.Cap, true)
		}
		rm.PublicRooms = rooms
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
				if (room.ClientCount == -1 && !room.Static) || time.Since(room.LastUpdate) > RoomTimeout {
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

func (rm *RoomManager) createRoom(roomName string, maxClients int, static bool) (*Room, error) {
	if len(rm.Rooms) < rm.MaxRooms {
		newRoomID := roomName
		if !static {
			newRoomID = rm.generateNewRoomID()
		} else {
			if strings.Contains(newRoomID, "-") {
				return nil, errors.New("static room names may not contain hyphens")
			}
		}
		newRoom := newRoom(rm, newRoomID, roomName, static, maxClients)
		rm.Rooms[newRoom.ID] = newRoom

		log.Println("[ROOM CREATED] - Created room ID \""+newRoomID+"\" | There are now", len(rm.Rooms), "active rooms.")
		return newRoom, nil
	}
	return nil, errors.New("the server is at maximum rooms capacity")
}

func (rm *RoomManager) closeRoom(roomID string) {
	if rm.Rooms[roomID].ClientCount > 0 {
		rm.Rooms[roomID].close()
	}
	delete(rm.Rooms, roomID)
	log.Println("[ROOM CLOSED] - Closed room ID \""+roomID+"\" | There are now", len(rm.Rooms), "active rooms.")
}
