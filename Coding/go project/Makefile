include .env
export

.PHONY: run build up test

run:
	@go run ./cmd/http.go

build:
	@go build -o ./bin/http ./cmd/http.go

up:
	@./bin/http
	
test:
	@go test ./modules/quiz/...