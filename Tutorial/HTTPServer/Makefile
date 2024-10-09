.DEFAULT_GOAL := build

.PHONY:fmt vet build run clean

fmt:
	go fmt ./...

vet: fmt	
	go vet ./...

build: vet
	go build

run: build
	go run ./...

clean:
	go clean ./...