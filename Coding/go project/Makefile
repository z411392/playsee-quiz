include .env
export

.PHONY: run build up

run:
	@go run ./main/http.go

build:
	@go build -o ./bin/http ./main/http.go

up:
	@./bin/http