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
		if !hasRoom {
			newRoom, err := rm.createRoom()
			if err != nil {
				log.Println("Failed to create room:", err)
				return
			}
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
