package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dio-av/http-service/app"
)

func runServer() {
	srv := app.NewServer(nil)

	server := &http.Server{
		Handler:      srv.Router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func main() {
	runServer()
}
