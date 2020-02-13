PROJECT = "go-lang"

# Need to load environment variables

ifeq ($(MAKECMDGOALS), start)
-include .env
endif

# To make sure not having conflicts with files defined the application
.PHONY: dependecies build start

# Dependecies
dependecies:
	dep ensure -v

# Build
build: dependecies
	go build -o ${PROJECT}

# Start
start: build
	./${PROJECT}

