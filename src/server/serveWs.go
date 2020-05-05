package server

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

//ServeWs serves a websocket to the client.
func (rm *RoomManager) ServeWs(w http.ResponseWriter, r *http.Request) {
	var err error
	var errCode int
	var client *client

	defer func() {
		if err != nil {
			log.Println("[JOIN FAIL] -", err)
			client.Cancel(errCode, err.Error())
			return
		}
	}()

	r.ParseForm()
	name, _ := r.Form["name"]
	roomID, hasRoom := r.Form["room"]

	//Ascertaining the client's IP and port from the initial request...
	sourceIPs := r.Header["X-Forwarded-For"]
	clientIP := r.RemoteAddr //Default to RemoteAddr so works on dev
	if len(sourceIPs) > 0 {
		port := strings.Split(clientIP, ":")[1]             //We need to keep the port
		clientIP = sourceIPs[len(sourceIPs)-1] + ":" + port //Client's actual IP is always the last one
	}

	client, err = newClient(w, r, name[0], clientIP)
	if err != nil {
		errCode = 4400
		return
	}

	if !hasRoom { //Client is trying to create a new room.
		newRoom := newRoom(rm, rm.generateNewRoomID(), "", DefaultRoomSize)
		err = rm.addRoom(newRoom)
		if err != nil {
			/*Current possible errors here:
			- The server has reached maximum rooms capacity.
			*/
			errCode = 4503
			return
		}

		/*As of 28/03/20 addClient should never return an error here as:
		  - The new room should be empty as it was just created
		  and
		  - addClient only returns an error if the room is already full.
		  However, this may change in future. So the error is still handled.
		*/
		client.room = newRoom
		err = newRoom.addClient(client)
		if err != nil {
			errCode = 4666
			return
		}

		log.Println("[JOIN SUCCESS] - Created room ID " + newRoom.getID() + "")
		client.GO()

	} else { //Client is attempting to join a room.
		room, roomExists := rm.Rooms[roomID[0]]
		if !roomExists {
			err = errors.New("this room doesn't exist (it may have closed)")
			errCode = 4404
			return
		}

		//Attempt to add client to the room (typically will fail if someone has already taken the name they're trying to join with)
		err = room.addClient(client)
		if err != nil {
			/*Current possible errors here:
			- The name the client wanted is already taken in this room
			- The room is already full
			*/
			errCode = 4409
			return
		}

		client.room = room
		log.Println("[JOIN SUCCESS] - Added client to room", roomID[0])
		client.GO()

	}
}
