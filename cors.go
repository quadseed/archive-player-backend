package main

import (
	"github.com/rs/cors"
	"net/http"
)

func CorsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8000",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowCredentials: true,
	})
}
