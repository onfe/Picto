package server

import (
	"log"
	"net/http"
)

//ServeWs serves a websocket to the client.
func (rm *RoomManager) ServeWs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name, hasName := r.Form["name"]
	roomID, hasRoom := r.Form["room"]

	if hasName {
		client, err := newClient(w, r, name[0])
		if err != nil {
			log.Println("Failed to create websocket:", err)
			return
		}

		if !hasRoom {
			newRoom, err := rm.createRoom()
			if err != nil {
				log.Println("Failed to create room:", err)
				client.closeConnection(err.Error())
				return
			}

			client.room = newRoom

			newRoom.addClient(client)
			log.Println("Created room", newRoom.ID, "for client with name:", client.Name)

		} else {
			if room, roomExists := rm.Rooms[roomID[0]]; roomExists {
				err = room.addClient(client)

				if err != nil {
					log.Println("Someone failed to join room ID"+roomID[0], "with name '"+client.Name+"':", err)
					client.closeConnection(err.Error())
					return
				}

				client.room = room
				log.Println("Added client '"+client.Name+"' (ID:"+client.ID+") to room", roomID[0])

			} else {
				log.Println("Client with name '" + name[0] + "' tried to join a room doesn't exist.")
				client.closeConnection("Room doesn't exist")
			}
		}
	}
}
