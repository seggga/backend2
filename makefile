PROJECT := github.com/seggga/backend2
GIT_COMMIT := $(shell git rev-parse HEAD)
VERSION := latest
APP_NAME := k8s-go-app

all: run

run:
	go build -ldflags="-X '$(PROJECT)/version.Version=$(VERSION)' -X '$(PROJECT)/version.Commit=$(GIT_COMMIT)'" -o $(APP_NAME) && ./$(APP_NAME)

