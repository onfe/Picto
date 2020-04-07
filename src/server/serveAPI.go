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
	method := "unset"
	var response []byte
	var err error

	//Setup the response.
	defer func() {
		if err != nil {
			//If there is an error, the error string becomes the response.
			log.Println("[API FAIL] - Method: " + method + ", Error: " + err.Error())
			response, _ = json.Marshal(err.Error())
		} else {
			log.Println("[API SUCCESS] - Method: " + method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}()

	//Parse the request.
	err = r.ParseForm()
	if err != nil {
		return
	}

	//First check if a method's actually been supplied.
	METHOD, methodSupplied := r.Form["method"]
	if !methodSupplied {
		err = errors.New("no method supplied")
		return
	}
	method = METHOD[0]

	//If a token is supplied, check the token supplied is valid
	token, tokenSupplied := r.Form["token"]
	if tokenSupplied && token[0] != rm.apiToken {
		//If it's not, we return here and now.
		err = errors.New("invalid token: " + token[0])
		return
	}

	//If we didn't just return and a token was supplied, it must be valid.
	if tokenSupplied {
		//Authenticated methods:
		switch method {

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
			roomID, roomIDSupplied := r.Form["id"]
			if !roomIDSupplied {
				err = errors.New("no `id` supplied")
				return
			}

			room, roomExists := rm.Rooms[roomID[0]]
			if !roomExists {
				err = errors.New("a room with that `id` does not exist")
				return
			}

			response, err = json.Marshal(room)
			return

		case "announce":
			message, messageSupplied := r.Form["message"]
			if !messageSupplied {
				err = errors.New("no `message` supplied")
				return
			}

			roomID, roomIDSupplied := r.Form["id"]
			if roomIDSupplied {
				if _, roomExists := rm.Rooms[roomID[0]]; !roomExists {
					err = errors.New("a room with that `id` does not exist")
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

			roomName, roomNameSupplied := r.Form["name"]
			if !roomNameSupplied {
				err = errors.New("no `name` supplied")
				return
			}

			_maxClients, maxClientsSupplied := r.Form["size"]
			if maxClientsSupplied {
				maxClients, err = strconv.Atoi(_maxClients[0])
				if err != nil {
					return
				}
				if maxClients < 1 {
					err = errors.New("`size` is too small (min size is 1)")
					return
				}
				if maxClients > MaxClientsPerRoom {
					err = errors.New("`size` is too big (max size is " + strconv.Itoa(MaxClientsPerRoom) + ")")
					return
				}
			}

			newRoom, ERR := rm.createRoom(roomName[0], maxClients, "static")
			err = ERR
			if err != nil {
				return
			}

			response, err = json.Marshal("new room created with `id` '" + newRoom.getID() + "'")
			return

		case "close_room":
			//default values
			reason := "This room is being closed by the server."
			closeTime := DefaultCloseTime

			roomID, roomIDSupplied := r.Form["id"]
			if !roomIDSupplied {
				err = errors.New("no `id` supplied")
				return
			}

			_reason, reasonSupplied := r.Form["reason"]
			if reasonSupplied {
				reason = _reason[0]
			}

			_closeTime, closeTimeSupplied := r.Form["time"]
			if closeTimeSupplied {
				closeTimeInt, ERR := strconv.Atoi(_closeTime[0])
				err = ERR
				if err != nil {
					return
				}
				if closeTimeInt < 0 {
					err = errors.New("`time` is too small (min time is 0)")
					return
				}
				closeTime = time.Duration(closeTimeInt) * time.Second
			}

			if _, roomExists := rm.Rooms[roomID[0]]; !roomExists {
				err = errors.New("a room with that `id` does not exist")
				return
			}

			rm.Rooms[roomID[0]].setCloseTime(time.Now().Add(closeTime))
			rm.Rooms[roomID[0]].announce(reason)
			rm.Rooms[roomID[0]].announce(fmt.Sprintf("Room closing in %.0f seconds...", closeTime.Seconds()))
			response, err = json.Marshal("closed room of `id` '" + roomID[0] + "'.")
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
					Name: r.Name,
					Cap:  r.ClientManager.MaxClients,
					Pop:  r.ClientManager.ClientCount,
				}
				i++
			}
			response, err = json.Marshal(roomStates)
			return

		default:
			err = errors.New("unrecognised method")
			return
		}
	} else { //If a token wasn't supplied, then we want public methods:
		//Public methods:
		switch method {

		case "room_exists":
			roomID, roomIDSupplied := r.Form["id"]

			if !roomIDSupplied {
				err = errors.New("no `id` supplied")
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
			for _, r := range rm.SubmissionRooms {
				if !r.Closing {
					roomStates = append(
						roomStates,
						roomState{
							Name: r.getID(),
							Cap:  r.ClientManager.MaxClients,
							Pop:  r.ClientManager.ClientCount,
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
