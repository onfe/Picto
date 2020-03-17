package server

import (
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
		hasName = !(name[0] == "")
	}

	if hasName {
		client, err := newClient(w, r, name[0])
		if err != nil {
			log.Println("[JOIN FAIL] - Failed to create websocket:", err)
			return
		}

		if !hasRoom { //Client is trying to create a new room.

			newRoom, err := rm.createRoom("Picto Room", false, DefaultRoomSize)
			if err != nil {
				log.Println("[JOIN FAIL] - Failed to create room:", err)
				client.closeConnection()
				return
			}

			client.room = newRoom
			newRoom.addClient(client)
			log.Println("[JOIN SUCCESS] - Created room \""+newRoom.ID+"\" for client with name:", client.Name)

		} else { //Client is attempting to join a room.
			if room, roomExists := rm.Rooms[roomID[0]]; roomExists {

				//Attempt to add client to the room (typically will fail if someone has already taken the name they're trying to join with)
				err = room.addClient(client)
				if err != nil {
					log.Println("[JOIN FAIL] - Someone failed to join room ID"+roomID[0], "with name '"+client.Name+"':", err)
					client.closeConnection()
					return
				}

				client.room = room
				log.Println("[JOIN SUCCESS] - Added client '"+client.Name+"' (ID:"+strconv.Itoa(client.ID)+") to room", roomID[0])

			} else { //If room doesn't exist...
				log.Println("[JOIN FAIL] - Client with name '" + name[0] + "' tried to join a room doesn't exist.")
				client.closeConnection()
			}
		}
	} else {
		if hasRoom {
			log.Println("[JOIN FAIL] - Client attempted to join room ID" + roomID[0] + " without a name.")
		} else {
			log.Println("[JOIN FAIL] - Client attempted to join without a name or room.")
		}
	}
}
