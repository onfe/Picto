package main

import (
	"log"
	"net/http"

	"github.com/onfe/Picto/src/server"
)

var roomManager server.RoomManager

func main() {
	roomManager = server.NewRoomManager(server.MaxRooms)

	fs := http.FileServer(http.Dir("../client/dist"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", roomManager.ServeWs)
	http.HandleFunc("/api/", roomManager.ServeAPI)

	log.Fatal(http.ListenAndServe(server.Address, nil))
}
