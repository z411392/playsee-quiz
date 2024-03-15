include .env
export

.PHONY: run build up

run:
	@go run ./cmd/http.go

build:
	@go build -o ./bin/http ./cmd/http.go

up:
	@./bin/http