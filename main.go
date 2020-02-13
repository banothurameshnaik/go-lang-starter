package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
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
}

func main() {
	// To read evnironment variables, we can use os package, as shown below
	goPath := os.Getenv("GOPATH")

	// To read environment variables from .env file, we need to have some external package
	logLevel, _ := os.LookupEnv("LOG_LEVEL")

	fmt.Println(fmt.Sprintf("Sample go Service || Environment: %s || Port: %s || GO Path: %s || Log Level: %s", *environment, *port, goPath, logLevel))
}
