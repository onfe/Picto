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
	Rooms           map[string]roomInterface `json:"Rooms"`
	MaxRooms        int                      `json:"MaxRooms"`
	apiToken        string
	Mode            string
	wordsList       []string
	StaticRooms     map[string]*staticRoom
	SubmissionRooms map[string]*submissionRoom
}

//NewRoomManager creates a new room manager.
func NewRoomManager(MaxRooms int, apiToken string, Mode string, wordsList []string) RoomManager {
	rm := RoomManager{
		Rooms:           make(map[string]roomInterface, MaxRooms),
		MaxRooms:        MaxRooms,
		apiToken:        apiToken,
		Mode:            Mode,
		wordsList:       wordsList,
		StaticRooms:     make(map[string]*staticRoom, MaxRooms),
		SubmissionRooms: make(map[string]*submissionRoom, MaxRooms),
	}
	go rm.roomMonitorLoop()
	return rm
}

/*LoadStaticRoomConfig loads a room config env var of static rooms.*/
func (rm *RoomManager) LoadStaticRoomConfig(varname string) error {
	config, configured := os.LookupEnv(varname)
	if !configured {
		return errors.New("couldn't find config var: " + varname)
	}

	type staticRoomConfig struct {
		Name string
		Cap  int
	}

	var roomConfigs []staticRoomConfig
	configBytes := []byte(config)

	err := json.Unmarshal(configBytes, &roomConfigs)
	if err != nil {
		return err
	}

	for _, r := range roomConfigs {
		newRoom := newStaticRoom(rm, r.Name, r.Cap)
		err = rm.addRoom(newRoom)
		if err != nil {
			return err
		}
	}

	return nil
}

/*LoadSubmissionRoomConfig loads a room config env var of static rooms.*/
func (rm *RoomManager) LoadSubmissionRoomConfig(varname string) error {
	config, configured := os.LookupEnv(varname)
	if !configured {
		return errors.New("couldn't find config var: " + varname)
	}

	type submissionRoomConfig struct {
		Name        string
		Description string
		Cap         int
	}

	var roomConfigs []submissionRoomConfig
	configBytes := []byte(config)

	err := json.Unmarshal(configBytes, &roomConfigs)
	if err != nil {
		return err
	}

	for _, r := range roomConfigs {
		newRoom := newSubmissionRoom(rm, r.Name, r.Description, r.Cap)
		err = rm.addRoom(newRoom)
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

func (rm *RoomManager) addRoom(newRoom roomInterface) error {
	if len(rm.Rooms) > rm.MaxRooms {
		return errors.New("the server is at maximum rooms capacity")
	}

	if newRoom.getType() != "dynamic" {
		/* We've got to check that non-dynamic room names don't contain hyphens to
		ensure there is no overlap between dynamic and other types of room. */
		if strings.Contains(newRoom.getID(), "-") {
			return errors.New("only dynamic room names may contain hyphens")
		}
	}

	/* We've also got to check the room name isn't alredy in use */
	for roomID := range rm.Rooms {
		if roomID == newRoom.getID() {
			return errors.New("a room already exists with that name")
		}
	}

	rm.Rooms[newRoom.getID()] = newRoom

	switch newRoom.getType() {
	case "static":
		rm.StaticRooms[newRoom.getID()] = newRoom.(*staticRoom)
	case "submission":
		rm.SubmissionRooms[newRoom.getID()] = newRoom.(*submissionRoom)
	}

	log.Println("[ROOM CREATED] - Created room ID \""+newRoom.getID()+"\" | There are now", len(rm.Rooms), "active rooms.")

	return nil
}

func (rm *RoomManager) closeRoom(roomID string) {
	rm.Rooms[roomID].close()

	if _, isStatic := rm.StaticRooms[roomID]; isStatic {
		delete(rm.StaticRooms, roomID)
	}

	delete(rm.Rooms, roomID)

	log.Println("[ROOM CLOSED] - Closed room ID \""+roomID+"\" | There are now", len(rm.Rooms), "active rooms.")
}
