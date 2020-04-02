package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/onfe/Picto/src/server"
)

var roomManager server.RoomManager

//CustomResponseWriter ...
type CustomResponseWriter struct {
	http.ResponseWriter
	status int
}

//WriteHeader ...
func (w *CustomResponseWriter) WriteHeader(status int) {
	w.status = status
	if status != http.StatusNotFound {
		w.ResponseWriter.WriteHeader(status)
	}
}

func (w *CustomResponseWriter) Write(data []byte) (int, error) {
	if w.status != http.StatusNotFound {
		return w.ResponseWriter.Write(data)
	}
	return len(data), nil
}

func getFileHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customResponseWriter := &CustomResponseWriter{ResponseWriter: w}

		h.ServeHTTP(customResponseWriter, r)

		if customResponseWriter.status == 404 {
			log.Println("[REDIRECT] - To index.html, from:", r.RequestURI)
			data, _ := ioutil.ReadFile("dist/index.html")
			w.Header().Set("Content-Type", "text/html")
			http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(data))
		}
	}
}

func loadWordsList(fp string) []string {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Println("[SYSTEM] - Couldn't open words list.")
		panic(err)
	}
	return strings.Split(strings.Replace(string(data), "\r\n", "\n", -1), "\n")
}

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
	wordsList := loadWordsList("words.txt")

	//Creating room manager instance
	apiToken, prod := os.LookupEnv("API_TOKEN") //in prod if API_TOKEN env variable is set.
	if prod {
		roomManager = server.NewRoomManager(server.MaxRooms, apiToken, "prod", wordsList, "PUBLIC_ROOMS")
	} else {
		roomManager = server.NewRoomManager(server.MaxRooms, "dev", "dev", wordsList, "PUBLIC_ROOMS")
	}

	//Seeing random number generator
	seedString, seeded := os.LookupEnv("RAND_SEED")
	if seeded {
		seed, err := strconv.ParseInt(seedString, 10, 64)
		if err != nil {
			log.Println("RAND_SEED set incorrectly (should be int64)")
		}
		if err == nil {
			rand.Seed(seed)
		}
	}

	//Setting up routing
	fs := getFileHandler(http.FileServer(http.Dir("dist")))
	http.Handle("/", fs)
	http.HandleFunc("/ws", roomManager.ServeWs)
	http.HandleFunc("/api/", roomManager.ServeAPI)

	address := ":8080"
	if prod {
		address = ":" + os.Getenv("PORT")
	}

	log.Fatal(http.ListenAndServe(address, nil))
}
