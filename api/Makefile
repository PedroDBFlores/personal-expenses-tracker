APP_NAME=api

.PHONY: all build run clean test tidy fmt lint

all: build

build:
	go build -ldflags="-s -w" -tags netgo -o $(APP_NAME)

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

test:
	go test ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...
