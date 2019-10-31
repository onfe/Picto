package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/onfe/Picto/src/server"
)

var roomManager server.RoomManager

func main() {
	_, prod := os.LookupEnv("API_TOKEN") //in prod if API_TOKEN env variable is set.

	roomManager = server.NewRoomManager(server.MaxRooms, prod)

	fs := http.FileServer(http.Dir("client/dist"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", roomManager.ServeWs)

	address := ":8080"
	if port, production := os.LookupEnv("PORT"); production {
		address = ":" + port
		if prod {
			http.HandleFunc("/api/", roomManager.ServeAPI)
		} else {
			http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				response, _ := json.Marshal("This feature has been disabled.")
				w.Write(response)
			})
		}
	} else {
		http.HandleFunc("/api/", roomManager.ServeAPI)
	}

	log.Fatal(http.ListenAndServe(address, nil))
}
