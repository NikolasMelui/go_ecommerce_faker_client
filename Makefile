.PHONY: build run

build:
	go build -v ./cmd/ecommerce_faker_client

run:
	go build -v ./cmd/ecommerce_faker_client; ./ecommerce_faker_client

.DEFAULT_GOAL := build
