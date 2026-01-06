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

gen-proto:
	mkdir -p proto/gen/study_service
    
	protoc --go_out=. --go_opt=module=github.com/Eklund2012/recall \
	    	--go-grpc_out=. --go-grpc_opt=module=github.com/Eklund2012/recall \
	    	proto/study.proto

	./services/brain/venv/Scripts/python -m grpc_tools.protoc -I./proto \
			--python_out=./services/brain \
			--grpc_python_out=./services/brain \
			proto/study.proto

.PHONY: all build run test clean
