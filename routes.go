package main

import (
	"github.com/gorilla/mux"
)

func handleRoutes(router *mux.Router) {
	router.HandleFunc("/", WelcomeHandler)
	router.HandleFunc("/healthcheck", HealthCheckHandler)
}
