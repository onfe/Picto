package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

//ServeAPI handles API calls.
func (rm *RoomManager) ServeAPI(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	method, methodSupplied := r.Form["method"]

	var response []byte
	var err error

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		if err != nil {
			log.Println("[API FAIL] - Method: " + method[0] + ", Error:" + err.Error())
			response, _ = json.Marshal(err.Error())
		} else {
			log.Println("[API SUCCESS] - Method: " + method[0])
		}
		w.Write(response)
	}()

	if !methodSupplied {
		err = errors.New("no method supplied")
		return
	}

	token, tokenSupplied := r.Form["token"]

	if tokenSupplied && token[0] != rm.apiToken {
		log.Println("[API FAIL] - An attempt to query the API was made with an invalid token:", token[0])
		response, err = json.Marshal("invalid token")
		return

	} else if tokenSupplied && token[0] == rm.apiToken {
		switch method[0] {

		case "get_state":
			if rm.Mode != "dev" {
				err = errors.New("this method is not available in prod")
				return
			}
			response, err = json.Marshal(rm)
			return

		case "get_mode":
			response, err = json.Marshal(rm.Mode)
			return

		case "get_room_ids":
			roomIDs := make([]string, 0, len(rm.Rooms))

			for roomID := range rm.Rooms {
				roomIDs = append(roomIDs, roomID)
			}

			response, err = json.Marshal(roomIDs)
			return

		case "get_room_state":
			roomID, roomIDSupplied := r.Form["room_id"]
			if !roomIDSupplied {
				err = errors.New("no room id supplied")
				return
			}

			room, roomExists := rm.Rooms[roomID[0]]
			if !roomExists {
				err = errors.New("room does not exist")
				return
			}

			response, err = json.Marshal(room)
			return

		case "announce":
			message, messageSupplied := r.Form["message"]
			roomID, roomIDSupplied := r.Form["room_id"]

			if !messageSupplied {
				err = errors.New("no message supplied")
				return
			}

			if roomIDSupplied {
				if _, roomExists := rm.Rooms[roomID[0]]; !roomExists {
					err = errors.New("room doesn't exist")
					return
				}
				rm.Rooms[roomID[0]].announce(message[0])
				response, err = json.Marshal("Announced '" + message[0] + "' to " + roomID[0])
				return
			}

			for _, room := range rm.Rooms {
				room.announce(message[0])
			}
			response, err = json.Marshal("Announced " + message[0] + " To all rooms")
			return

		case "create_static_room":
			//Default values
			maxClients := DefaultRoomSize
			public := false

			roomName, roomNameSupplied := r.Form["room_name"]
			if !roomNameSupplied {
				err = errors.New("no room name supplied")
				return
			}

			_public, publicSupplied := r.Form["public"]
			if publicSupplied {
				public = _public[0] == "true"
			}

			_maxClients, maxClientsSupplied := r.Form["room_size"]
			if maxClientsSupplied {
				maxClients, err = strconv.Atoi(_maxClients[0])
				if err != nil {
					err = errors.New("size must be an integer value")
					return
				}
				if maxClients < 1 {
					err = errors.New("size is too small (min size is 1)")
					return
				}
				if maxClients > MaxClientsPerRoom {
					err = errors.New("size is too big (max size is " + strconv.Itoa(MaxClientsPerRoom) + ")")
					return
				}
			}

			if _, roomExists := rm.Rooms[roomName[0]]; roomExists {
				err = errors.New("a room already exists with that name")
				return
			}

			newRoom, err := rm.createRoom(roomName[0], maxClients, true, public)
			if err != nil {
				return
			}

			response, err = json.Marshal("new room created with id '" + newRoom.ID + "'")
			return

		case "close_room":
			//default values
			reason := "This room is being closed by the server."

			roomID, roomIDSupplied := r.Form["room_id"]
			if !roomIDSupplied {
				err = errors.New("no id supplied")
				return
			}

			_reason, reasonSupplied := r.Form["reason"]
			if reasonSupplied {
				reason = _reason[0]
			}

			closeTime := 10 //seconds
			_closeTime, closeTimeSupplied := r.Form["close_time"]
			if closeTimeSupplied {
				closeTime, err = strconv.Atoi(_closeTime[0])
				if err != nil {
					return
				}
			}

			if _, roomExists := rm.Rooms[roomID[0]]; !roomExists {
				err = errors.New("room doesn't exist")
				return
			}

			if rm.Rooms[roomID[0]].Closing {
				err = errors.New("room is already closing")
				return
			}

			rm.Rooms[roomID[0]].Closing = true
			rm.Rooms[roomID[0]].announce(reason)
			rm.Rooms[roomID[0]].announce(fmt.Sprintf("Room closing in %d seconds...", closeTime))
			go func(rm *RoomManager) {
				time.Sleep(time.Duration(closeTime) * time.Second)
				rm.closeRoom(roomID[0])
			}(rm)
			response, err = json.Marshal("closed room of id '" + roomID[0] + "'.")
			return

		case "get_static_rooms":
			type roomState struct {
				Name   string
				Public bool
				Cap    int
				Pop    int
			}
			roomStates := make([]roomState, len(rm.StaticRooms))
			i := 0
			for _, r := range rm.StaticRooms {
				roomStates[i] = roomState{
					Name:   r.Name,
					Public: r.Public,
					Cap:    r.MaxClients,
					Pop:    r.ClientCount,
				}
				i++
			}
			response, err = json.Marshal(roomStates)
			return

		default:
			err = errors.New("unrecognised method")
			return
		}
	} else {
		switch method[0] {

		case "room_exists":
			roomID, roomIDSupplied := r.Form["room_id"]

			if !roomIDSupplied {
				err = errors.New("no id supplied")
				return
			}

			_, hasRoom := rm.Rooms[roomID[0]]
			response, err = json.Marshal(hasRoom)
			return

		case "get_public_rooms":
			type roomState struct {
				Name string
				Cap  int
				Pop  int
			}
			var roomStates []roomState
			for _, r := range rm.StaticRooms {
				if r.Public && !r.Closing {
					roomStates = append(
						roomStates,
						roomState{
							Name: r.Name,
							Cap:  r.MaxClients,
							Pop:  r.ClientCount,
						})
				}
			}

			sort.Slice(roomStates[:], func(i, j int) bool {
				if roomStates[i].Pop != roomStates[j].Pop {
					//Populations sorted highest first
					return roomStates[i].Pop > roomStates[j].Pop
				} else if roomStates[i].Cap != roomStates[j].Cap {
					//Caps sorted highest first
					return roomStates[i].Cap > roomStates[j].Cap
				} else {
					//Names sorted A-Z
					return roomStates[i].Name[0] < roomStates[j].Name[0]
				}
			})

			response, err = json.Marshal(roomStates)
			return

		default:
			err = errors.New("unrecogised method")
			return
		}
	}
}
