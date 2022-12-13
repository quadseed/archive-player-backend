package main

import (
	"github.com/quadseed/archive-player-backend/handler"
	"log"
	"net/http"
)

func init() {
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.RootHandler)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
