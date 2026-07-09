APP_NAME := goman
MAIN_PKG := ./cmd/api
BIN_DIR := bin
BIN := $(BIN_DIR)/$(APP_NAME)
AIR := air
COMPOSE ?= podman compose
COMPOSE_FILE := docker-compose.yml
COMPOSE_WATCH_FILE := docker-compose.watch.yml

.PHONY: all build run watch air-install test fmt tidy clean docker-build docker-up docker-down docker-logs docker-ps postgres-up postgres-down docker-watch-build docker-watch-up docker-watch-down docker-watch-logs

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

docker-build:
	$(COMPOSE) -f $(COMPOSE_FILE) build

docker-up:
	$(COMPOSE) -f $(COMPOSE_FILE) up -d --build

docker-watch-build:
	$(COMPOSE) -f $(COMPOSE_FILE) -f $(COMPOSE_WATCH_FILE) build

docker-watch-up:
	$(COMPOSE) -f $(COMPOSE_FILE) -f $(COMPOSE_WATCH_FILE) up -d

docker-down:
	$(COMPOSE) -f $(COMPOSE_FILE) down

docker-watch-down:
	$(COMPOSE) -f $(COMPOSE_FILE) -f $(COMPOSE_WATCH_FILE) down

docker-logs:
	$(COMPOSE) -f $(COMPOSE_FILE) logs -f

docker-watch-logs:
	$(COMPOSE) -f $(COMPOSE_FILE) -f $(COMPOSE_WATCH_FILE) logs -f

docker-ps:
	$(COMPOSE) -f $(COMPOSE_FILE) ps

postgres-up:
	$(COMPOSE) -f $(COMPOSE_FILE) up -d postgres

postgres-down:
	$(COMPOSE) -f $(COMPOSE_FILE) stop postgres
