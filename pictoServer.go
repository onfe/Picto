package main

import (
	"log"
	"net/http"
	"os"

	"github.com/onfe/Picto/src/server"
)

var roomManager server.RoomManager

func main() {
	apiToken, prod := os.LookupEnv("API_TOKEN") //in prod if API_TOKEN env variable is set.
	if prod {
		roomManager = server.NewRoomManager(server.MaxRooms, apiToken, "prod")
	} else {
		roomManager = server.NewRoomManager(server.MaxRooms, "dev", "dev")
	}

	fs := http.FileServer(http.Dir("client/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", roomManager.ServeWs)
	http.HandleFunc("/api/", roomManager.ServeAPI)

	address := ":8080"
	if prod {
		address = ":" + os.Getenv("PORT")
	}

	log.Fatal(http.ListenAndServe(address, nil))
}
