USERNAME := seggga
APP_NAME := static-srv
VERSION := 1.0.2

#write here path for your project
PROJECT := github.com/seggga/backend2
GIT_COMMIT := $(shell git rev-parse HEAD)


all: run

run:
	go build -ldflags="-X '$(PROJECT)/version.Version=$(VERSION)' \
	-X '$(PROJECT)/version.Commit=$(GIT_COMMIT)'" -o $(APP_NAME) && ./$(APP_NAME)

build_container:
	docker build --build-arg GIT_COMMIT=$(GIT_COMMIT) --build-arg VERSION=$(VERSION)  --build-arg=PROJECT=$(PROJECT) \
	-t docker.io/$(USERNAME)/$(APP_NAME):$(VERSION) .


push_container:
	docker push  docker.io/$(USERNAME)/$(APP_NAME):$(VERSION)