package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WelcomeHandler Welcome endpoint handler
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to sample go-lang application!!!")
}

// HealthCheckHandler Health check end point handler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"OK! I am healthy!"})
}
