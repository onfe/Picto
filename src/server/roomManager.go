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
func NewRoomManager(MaxRooms int, apiToken string, Mode string, wordsList []string) RoomManager {
	rm := RoomManager{
		Rooms:           make(map[string]RoomInterface, MaxRooms),
		MaxRooms:        MaxRooms,
		apiToken:        apiToken,
		Mode:            Mode,
		wordsList:       wordsList,
		StaticRooms:     make(map[string]*StaticRoom, MaxRooms),
		SubmissionRooms: make(map[string]*SubmissionRoom, MaxRooms),
	}
	go rm.roomMonitorLoop()
	return rm
}

/*LoadRoomConfig loads a room config env var of rooms of a given type.*/
func (rm *RoomManager) LoadRoomConfig(varname, roomType string) error {
	log.Println(varname, roomType)
	config, configured := os.LookupEnv(varname)
	log.Println(config, configured)
	if !configured {
		return errors.New("couldn't find config var: " + varname)
	}

	type roomConfig struct {
		Name string
		Cap  int
	}

	var roomConfigs []roomConfig
	configBytes := []byte(config)

	err := json.Unmarshal(configBytes, &roomConfigs)
	if err != nil {
		return err
	}

	for _, r := range roomConfigs {
		_, err = rm.createRoom(r.Name, r.Cap, roomType)
		if err != nil {
			return err
		}
	}

	return nil
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

func (rm *RoomManager) createRoom(newRoomName string, maxClients int, roomType string) (RoomInterface, error) {
	if len(rm.Rooms) > rm.MaxRooms {
		return nil, errors.New("the server is at maximum rooms capacity")
	}

	var room RoomInterface

	if roomType != "dynamic" {
		/* We've got to check that static room names don't contain hyphens to
		ensure there is no overlap between dynamic and other types of room. */
		if strings.Contains(newRoomName, "-") {
			return nil, errors.New("only dynamic room names may contain hyphens")
		}
	} else {
		/* We've also got to check the room name isn't alredy in use */
		for roomID := range rm.Rooms {
			if roomID == newRoomName {
				return nil, errors.New("a room already exists with that name")
			}
		}
	}

	switch roomType {
	case "static":
		room = newStaticRoom(rm, newRoomName, maxClients)
		rm.StaticRooms[room.getID()] = room.(*StaticRoom)

	case "submission":
		log.Println(newRoomName)
		room = newSubmissionRoom(rm, newRoomName, maxClients)
		rm.SubmissionRooms[room.getID()] = room.(*SubmissionRoom)

	case "dynamic":
		room = newRoom(rm, rm.generateNewRoomID(), newRoomName, maxClients)

	default:
		return nil, errors.New("unrecognised room type")
	}

	rm.Rooms[room.getID()] = room

	log.Println("[ROOM CREATED] - Created room ID \""+room.getID()+"\" | There are now", len(rm.Rooms), "active rooms.")

	return room, nil
}

func (rm *RoomManager) closeRoom(roomID string) {
	rm.Rooms[roomID].close()

	if _, isStatic := rm.StaticRooms[roomID]; isStatic {
		delete(rm.StaticRooms, roomID)
	}

	delete(rm.Rooms, roomID)

	log.Println("[ROOM CLOSED] - Closed room ID \""+roomID+"\" | There are now", len(rm.Rooms), "active rooms.")
}
