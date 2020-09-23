package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/onfe/Picto/src/server"
)

var roomManager server.RoomManager

func main() {
	//Logfile setup
	logFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	//Starting performance monitor
	go NewMonitor(60)

	//Loading words list
	wordsList := server.LoadWordsList("words.txt")

	//Getting env var values
	apiToken, tokenSet := os.LookupEnv("API_TOKEN")
	if !tokenSet {
		log.Println("[ENV VAR] - API_TOKEN not set, defaulting to 'dev'")
	}
	port, portSet := os.LookupEnv("PORT")
	if !portSet {
		log.Println("[ENV VAR] - PORT not set, defaulting to 8080")
		port = "8080"
	}

	//Creating room manager instance
	if tokenSet && portSet { //Only in prod if both API_TOKEN and PORT env variables are set.
		log.Println("[ROOM MANAGER] - Creating room manager in PROD mode")
		roomManager = server.NewRoomManager(server.MaxRooms, apiToken, "prod", wordsList)
	} else {
		log.Println("[ROOM MANAGER] - Creating room manager in DEV mode")
		roomManager = server.NewRoomManager(server.MaxRooms, "dev", "dev", wordsList)
	}

	err = roomManager.LoadStaticRoomConfig("STATIC_ROOMS")
	if err != nil {
		log.Println("Error loading STATIC_ROOMS:", err.Error())
	}

	err = roomManager.LoadModeratedRoomConfig("MODERATED_ROOMS")
	if err != nil {
		log.Println("Error loading MODERATED_ROOMS:", err.Error())
	}

	//Seeing random number generator
	rand.Seed(time.Now().UnixNano() / int64(time.Millisecond))

	//Setting up routing
	fs := getFileHandler(http.FileServer(http.Dir("dist")))
	http.Handle("/", fs)
	http.HandleFunc("/ws", roomManager.ServeWs)
	http.HandleFunc("/api/", roomManager.ServeAPI)
	log.Println("Serving on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
