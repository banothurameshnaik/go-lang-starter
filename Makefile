PROJECT = "go-lang"

.PHONY: dependecies 

# Dependecies
dependecies:
	dep ensure -v

# Build
build: dependecies
	go build -o ${PROJECT}

# Start
start: build
	./${PROJECT}

