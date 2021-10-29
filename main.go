package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
)

var (
	environment = flag.String("environment", "local", "Application Environment")
	port        = flag.String("port", "1255", "Application port number")
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	// Setting log formatter
	logger.SetFormatter(&logger.JSONFormatter{})
}

func main() {
	// Parsing flags
	flag.Parse()
	// To read evnironment variables, we can use os package, as shown below
	goPath := os.Getenv("GOPATH")

	// To read environment variables from .env file, we need to have some external package
	logLevel, _ := os.LookupEnv("LOG_LEVEL")

	// Advanced logging using logrus package, setting log level
	loggerLevel, _ := logger.ParseLevel(logLevel)
	logger.SetLevel(loggerLevel)

	logger.Info(fmt.Sprintf("Sample go Service || Environment: %s || Port: %s || GO Path: %s || Log Level: %s", *environment, *port, goPath, logLevel))

	// Logging
	// Simple logging using log go package
	log.Println("Simple logging looks like this")

	// Advanced logging using logrus package
	logger.Info("Advaced logging looks like this")

	// HTTP Server
	r := mux.NewRouter()
	handleRoutes(r)
	http.Handle("/", r)

	logger.Info(fmt.Sprintf("Starting the application and listing on port %s", *port))
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)

}
