package server

import (
	"encoding/json"
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

	if token, tokenSupplied := r.Form["token"]; tokenSupplied && token[0] == rm.apiToken {
		if !methodSupplied {
			log.Println("[API FAIL] - An attempt to query the API was made without supplying a method with token:", token[0])
			return
		}

		var response []byte
		var err error

		switch method[0] {

		case "get_state":
			if rm.Mode == "dev" {
				response, err = json.Marshal(rm)
			} else {
				response, err = json.Marshal("This method is not available in prod.")
			}

		case "get_mode":
			response, err = json.Marshal(rm.Mode)

		case "get_room_ids":
			roomIDs := make([]string, 0, len(rm.Rooms))
			for roomID := range rm.Rooms {
				roomIDs = append(roomIDs, roomID)
			}
			response, err = json.Marshal(roomIDs)

		case "get_room_state":
			roomID, roomIDSupplied := r.Form["room_id"]
			if !roomIDSupplied {
				response, err = json.Marshal("No room ID supplied.")
			} else {
				room, roomExists := rm.Rooms[roomID[0]]
				if !roomExists {
					response, err = json.Marshal("Room does not exist.")
				} else {
					response, err = json.Marshal(room)
				}
			}

		case "announce":
			message, messageSupplied := r.Form["message"]
			roomID, roomIDSupplied := r.Form["room_id"]
			if messageSupplied {
				if roomIDSupplied {
					rm.Rooms[roomID[0]].announce(message[0])
					response, err = json.Marshal("Announced " + message[0] + " To room ID" + roomID[0])
				} else {
					for _, room := range rm.Rooms {
						room.announce(message[0])
					}
					response, err = json.Marshal("Announced " + message[0] + " To all rooms")
				}
			} else {
				response, err = json.Marshal("Malformed API call. Please supply a message.")
			}

		case "create_static_room":
			roomName := ""
			maxClients := DefaultRoomSize
			public := false

			_roomName, roomNameSupplied := r.Form["room_name"]
			_maxClients, maxClientsSupplied := r.Form["room_size"]

			_public, publicSupplied := r.Form["public"]
			if publicSupplied {
				public = _public[0] == "true"
			}

			if !roomNameSupplied {
				response, err = json.Marshal("a room name must be supplied")
			} else {
				roomName = _roomName[0]

				if maxClientsSupplied {
					maxClients, err = strconv.Atoi(_maxClients[0])
				}

				if err != nil {
					response, err = json.Marshal("size supplied couldn't be converted to an integer value: " + err.Error())
				} else {
					if maxClients > MaxClientsPerRoom {
						response, err = json.Marshal("the max clients per room is " + strconv.Itoa(MaxClientsPerRoom))
					} else {
						newRoom, err := rm.createRoom(roomName, maxClients, true, public)
						if err != nil {
							response, err = json.Marshal("New room couldn't be created: " + err.Error())
						} else {
							response, err = json.Marshal("new room created with id '" + newRoom.ID + "'")
						}
					}
				}

			}

		case "close_room":
			roomID, roomIDSupplied := r.Form["room_id"]

			reason := "This room is being closed by the server."
			_reason, reasonSupplied := r.Form["reason"]
			if reasonSupplied {
				reason = _reason[0]
			}

			closeTime := 10 //seconds
			_closeTime, closeTimeSupplied := r.Form["close_time"]
			if closeTimeSupplied {
				closeTime, err = strconv.Atoi(_closeTime[0])
			}

			if err != nil {
				response, err = json.Marshal("Malformed API call. close_time must be an integer value")
			} else {
				if roomIDSupplied {
					if _, roomExists := rm.Rooms[roomID[0]]; roomExists {
						if !rm.Rooms[roomID[0]].Closing {
							rm.Rooms[roomID[0]].Closing = true
							rm.Rooms[roomID[0]].announce(reason)
							rm.Rooms[roomID[0]].announce(fmt.Sprintf("Room closing in %d seconds...", closeTime))
							go func(rm *RoomManager) {
								time.Sleep(time.Duration(closeTime) * time.Second)
								rm.closeRoom(roomID[0])
							}(rm)
							response, err = json.Marshal("closed room of id '" + roomID[0] + "'.")
						} else {
							response, err = json.Marshal("room is already closing.")
						}
					} else {
						response, err = json.Marshal("room_id supplied doesn't exist.")
					}
				} else {
					response, err = json.Marshal("Malformed API call. Please supply a room_id.")
				}
			}

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
			response, _ = json.Marshal(roomStates)

		default:
			response, err = json.Marshal("Unrecognised API method")

		}

		if err != nil {
			response, _ = json.Marshal(err)
		}

		log.Println("[PRIVATE API SUCCESS] - Method: " + method[0] + ", Result: " + string(response))
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	} else if !tokenSupplied && methodSupplied {
		var response []byte
		var err error

		switch method[0] {
		case "room_exists":
			roomID, roomIDSupplied := r.Form["room_id"]

			if roomIDSupplied {
				_, hasRoom := rm.Rooms[roomID[0]]
				response, err = json.Marshal(hasRoom)
			} else {
				response, err = json.Marshal("Malformed API call. Please supply a room_id.")
			}

			if err != nil {
				response, _ = json.Marshal(err)
			}

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
			response, _ = json.Marshal(roomStates)

		default:
			response, err = json.Marshal("Unrecognised API method")
		}

		if err != nil {
			response, _ = json.Marshal(err)
		}

		log.Println("[PUBLIC API SUCCESS] - Method: " + method[0] + ", Result: " + string(response))
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	} else if tokenSupplied {
		log.Println("[API FAIL] - An attempt to query the API was made with an invalid token:", token[0])
	} else {
		log.Println("[API FAIL] - An attempt to query the API was made without a token.")
	}
}
