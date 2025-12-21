SHELL := /usr/bin/bash

all: build test

build:
	go build -o recall.exe main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -f recall.exe

.PHONY: all build run test clean
