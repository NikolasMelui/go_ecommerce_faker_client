.PHONY: build run

build:
	go build -v ./cmd/efclient

run:
	go build -v ./cmd/efclient; ./efclient

.DEFAULT_GOAL := build
