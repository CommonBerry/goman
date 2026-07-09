APP_NAME := goman
MAIN_PKG := ./cmd/api
BIN_DIR := bin
BIN := $(BIN_DIR)/$(APP_NAME)
AIR := air

.PHONY: all build run watch air-install test fmt tidy clean

all: build

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

build: $(BIN_DIR)
	go build -o $(BIN) $(MAIN_PKG)

run:
	go run $(MAIN_PKG)

watch:
	$(AIR)

air-install:
	go install github.com/air-verse/air@latest

test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)
