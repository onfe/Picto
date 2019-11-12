package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/onfe/Picto/src/server"
)

var roomManager server.RoomManager

type CustomResponseWriter struct {
	http.ResponseWriter
	status int
}

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
			log.Printf("Redirecting %s to index.html.", r.RequestURI)
			data, _ := ioutil.ReadFile("client/dist/index.html")
			w.Header().Set("Content-Type", "text/html")
			http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(data))
		}
	}
}

func main() {
	apiToken, prod := os.LookupEnv("API_TOKEN") //in prod if API_TOKEN env variable is set.
	if prod {
		roomManager = server.NewRoomManager(server.MaxRooms, apiToken, "prod")
	} else {
		roomManager = server.NewRoomManager(server.MaxRooms, "dev", "dev")
	}

	fs := getFileHandler(http.FileServer(http.Dir("client/dist")))
	http.Handle("/", fs)
	http.HandleFunc("/ws", roomManager.ServeWs)
	http.HandleFunc("/api/", roomManager.ServeAPI)

	address := ":8080"
	if prod {
		address = ":" + os.Getenv("PORT")
	}

	log.Fatal(http.ListenAndServe(address, nil))
}