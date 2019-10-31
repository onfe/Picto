package main

import (
	"log"
	"net/http"
	"server"
)

var roomManager server.RoomManager

func serveHomepage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	}
}

func main() {
	roomManager = server.NewRoomManager(server.MaxRooms)

	http.HandleFunc("/", serveHomepage)
	http.HandleFunc("/ws", roomManager.ServeWs)
	http.HandleFunc("/api/", roomManager.ServeAPI)

	log.Fatal(http.ListenAndServe(server.Address, nil))
}
