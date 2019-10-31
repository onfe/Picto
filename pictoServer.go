package main

import (
	"log"
	"net/http"
	"os"

	"github.com/onfe/Picto/src/server"
)

var roomManager server.RoomManager

func main() {
	roomManager = server.NewRoomManager(server.MaxRooms)

	fs := http.FileServer(http.Dir("../client/dist"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", roomManager.ServeWs)
	http.HandleFunc("/api/", roomManager.ServeAPI)

	address := ":8080"
	if os.Getenv("GO_ENV") == "PRODUCTION" {
		address = ":80"
	}

	log.Fatal(http.ListenAndServe(address, nil))
}
