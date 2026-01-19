.PHONY: build run test fmt vet tidy lint clean help

BINARY_NAME=memory-map
BUILD_DIR=bin

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server

# run: build
# 	./$(BUILD_DIR)/$(BINARY_NAME)
run:
	go run ./cmd/server/main.go

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

clean:
	rm -rf $(BUILD_DIR)

help:
	@echo "Available targets: build run test fmt vet tidy lint clean"

## migration: Create a new migration file (usage: make migration name=add_tags)
migration:
	migrate create -ext sql -dir migrations -seq $(name)
