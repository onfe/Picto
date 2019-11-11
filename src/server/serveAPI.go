package server

import (
	"encoding/json"
	"log"
	"net/http"
)

//ServeAPI handles API calls.
func (rm *RoomManager) ServeAPI(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if token, tokenSupplied := r.Form["token"]; tokenSupplied && token[0] == rm.apiToken {
		method, methodSupplied := r.Form["method"]
		if !methodSupplied {
			log.Println("An attempt to query the API was made without supplying a method with token:", token[0])
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
			roomID, roomIDSupplied := r.Form["roomid"]
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

		default:
			response, err = json.Marshal("Unrecognised API method")

		}

		if err != nil {
			response, _ = json.Marshal(err)
		}

		log.Println(method[0]+":", string(response))
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	} else if tokenSupplied {
		log.Println("An attempt to query the API was made with an invalid token:", token[0])
	} else {
		log.Println("An attempt to query the API was made without a token.")
	}
}
