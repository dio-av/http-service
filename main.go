package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func runServer() {
	router := mux.NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func main() {
	runServer()
}
