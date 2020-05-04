package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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
