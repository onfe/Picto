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
		w.Header().Set("Content-Type", "application/json")

		method, _ := r.Form["method"]

		var response []byte
		var err error

		switch method[0] {
		case "get_state":
			if !rm.prod {
				response, err = json.Marshal(rm)
			} else {
				response, err = json.Marshal("This method is not available in prod.")
			}
		case "get_room_ids":
			roomIDs := make([]string, 0, len(rm.Rooms))
			for roomID, _ := range rm.Rooms {
				roomIDs = append(roomIDs, roomID)
			}
			response, err = json.Marshal(roomIDs)
		default:
			response, err = json.Marshal("Unrecognised API method")
		}

		if err != nil {
			response, _ = json.Marshal(err)
		}

		log.Println(method[0]+":", string(response))
		w.Write(response)

	} else {
		log.Println("An attempt to query the API was made with an invalid token:", token[0])
	}
}
