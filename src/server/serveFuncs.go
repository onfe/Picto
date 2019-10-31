package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//ServeWs serves a websocket to the client.
func (rm *RoomManager) ServeWs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name, hasName := r.Form["name"]
	roomID, hasRoom := r.Form["room"]
	if hasName {
		if !hasRoom {
			newRoom := rm.createRoom()
			client, err := newClient(w, r, newRoom, 0, name[0])
			if err != nil {
				log.Println("Failed to create websocket:", err)
				return
			}
			newRoom.addClient(client)
			log.Println("Created room", newRoom.ID, "for client with name:", client.Name)
		} else {
			if room, roomExists := rm.Rooms[roomID[0]]; roomExists {
				client, err := newClient(w, r, room, len(room.Clients), name[0])
				if err != nil {
					log.Println("Failed to create websocket:", err)
					return
				}
				nameFreeInRoom := room.addClient(client)
				if !nameFreeInRoom {
					log.Println("Someone tried to join room ID"+roomID[0], "with name '"+client.Name+"' but someone already claimed it.")
				} else {
					log.Println("Added client '"+client.Name+"' (ID"+strconv.Itoa(client.ID)+") to room", roomID[0])
				}
			} else {
				log.Println("Client with name '" + name[0] + "' tried to join a room doesn't exist.")
			}
		}
	}
}

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
			response, err = json.Marshal(rm)
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
