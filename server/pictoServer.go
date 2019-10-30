package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	//Address = Server address.
	Address = ":8080"
	//MaxRooms = maximum amount of rooms the server may have at any one time
	MaxRooms = 10
	//MaxRoomSize = Max size of default room.
	MaxRoomSize = 8
	//MaxMessageSize = Max size of a message from the client.
	MaxMessageSize = 1024
	//MinMessageInterval = Minimum interval between messages sent by a client to be acknowledged.
	MinMessageInterval = time.Second
	//ChatHistoryLen = Number of messages kept by server per room.
	ChatHistoryLen = 10
	//ClientSendTimeout is the time allotted for a message to be sent.
	ClientSendTimeout = 10 * time.Second
	//ClientTimeout = Max interval allotted between pings and pongs.
	ClientTimeout = 60 * time.Second
	//ClientPingPeriod is the period between pings.
	ClientPingPeriod = ClientTimeout / 10
)

func serveHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name, hasName := r.Form["name"]
	room, hasRoom := r.Form["room"]
	if hasName {
		if !hasRoom {
			newRoom := roomManager.createRoom()
			client, err := newClient(w, r, newRoom, 0, name[0])
			if err != nil {
				fmt.Println("Failed to create websocket:", err)
				return
			}
			newRoom.addClient(client)
			fmt.Println("Created room", newRoom.id)
		} else {
			if room, roomExists := roomManager.rooms[room[0]]; roomExists {
				client, err := newClient(w, r, &room, 0, name[0])
				if err != nil {
					fmt.Println("Failed to create websocket:", err)
					return
				}
				room.addClient(client)
				fmt.Println("Added client to room", room.id)
			} else {
				fmt.Println("Room doesn't exist.")
			}
		}
	}
}

var roomManager RoomManager

func main() {
	roomManager = newRoomManager(MaxRooms)

	http.HandleFunc("/", serveHomepage)
	http.HandleFunc("/api/ws", serveWs)

	http.ListenAndServe(Address, nil)
}
