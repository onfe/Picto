package server

import (
	"log"
	"net/http"
	"strconv"
)

//ServeWs serves a websocket to the client.
func (rm *RoomManager) ServeWs(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	name, _ := r.Form["name"]
	roomID, hasRoom := r.Form["room"]

	client, err := newClient(w, r, name[0])
	if err != nil {
		log.Println("[JOIN FAIL] - Failed to create websocket:", err)
		return
	}

	if err == nil {
		if !hasRoom { //Client is trying to create a new room.
			newRoom, err := rm.createRoom("Picto Room", false, DefaultRoomSize)
			if err != nil {
				/*Current possible errors here:
				- The server has reached maximum rooms capacity.
				*/
				log.Println("[JOIN FAIL] - Failed to create room:", err)
				client.Cancel(4001, err.Error())
				return
			}

			client.room = newRoom
			newRoom.addClient(client)
			log.Println("[JOIN SUCCESS] - Created room \""+newRoom.ID+"\" for client with name:", client.Name)
			client.GO()

		} else { //Client is attempting to join a room.
			if room, roomExists := rm.Rooms[roomID[0]]; roomExists {

				//Attempt to add client to the room (typically will fail if someone has already taken the name they're trying to join with)
				err = room.addClient(client)
				if err != nil {
					/*Current possible errors here:
					- The name the client wanted is already taken in this room
					- The room is already full
					*/
					log.Println("[JOIN FAIL] - Someone failed to join room ID"+roomID[0], "with name '"+client.Name+"':", err)
					client.Cancel(4001, err.Error())
					return
				}

				client.room = room
				log.Println("[JOIN SUCCESS] - Added client '"+client.Name+"' (ID:"+strconv.Itoa(client.ID)+") to room", roomID[0])
				client.GO()

			} else { //If room doesn't exist...
				log.Println("[JOIN FAIL] - Client with name '" + name[0] + "' tried to join a room doesn't exist.")
				client.Cancel(4404, "that room doesn't exist")
			}
		}
	}
}
