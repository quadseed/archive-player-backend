package main

import "github.com/rs/cors"

func CorsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8000",
			"https://localhost:8000",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"DELETE",
		},
		AllowCredentials: true,
	})
}
