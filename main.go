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

	corsConfig := CorsConfig()
	handler := corsConfig.Handler(mux)

	server := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}

	log.Fatal(server.ListenAndServe())
}
