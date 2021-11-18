PROJECT = "go-lang"

# Need to load environment variables

ifeq ($(MAKECMDGOALS), start)
-include .env
endif

ifeq ($(MAKECMDGOALS), local)
-include .env
endif

# To make sure not having conflicts with files defined the application
.PHONY: dependecies build start local stop

# Dependecies
dependecies:
	go mod vendor

# Build
build: dependecies
	go build -o $(PROJECT) .

# Build for linux
build-linux: dependecies
	GODEBUG=x509ignoreCN=0 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(PROJECT) .

# Start
start: build
	./${PROJECT} --environment=${ENVIRONMENT} --port=${PORT}

# Stop Commands for stopping docker
stop:
	docker container stop $(PROJECT) || echo "Container $(PROJECT) is alreay stopped"
	docker container rm $(PROJECT) || echo "Container $(PROJECT) is alreay removed"

# Running local with docker
local: build-linux stop
	docker build -t ${PROJECT} .
	docker container run -d -p ${PORT}:${PORT} \
	-e ENVIRONMENT=${ENVIRONMENT} \
	-e PORT=${PORT} \
	--name ${PROJECT} ${PROJECT}
