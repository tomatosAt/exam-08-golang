.PHONY: all fmt vet tidy build run test clean docker-up docker-down

BINARY := exam-08-golang

all: build

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

build:
	mkdir -p bin
	go build -v -o bin/$(BINARY) ./cmd/server

run:
	go run ./cmd/server start

test:
	go test ./...

clean:
	rm -rf bin

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
