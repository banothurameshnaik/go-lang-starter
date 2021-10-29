package main

import (
	"github.com/gorilla/mux"
)

func handleRoutes(router *mux.Router) {
	router.HandleFunc("/", WelcomeHandler).Methods("GET")
	router.HandleFunc("/healthcheck", HealthCheckHandler).Methods("GET")
}
