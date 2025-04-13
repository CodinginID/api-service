APP_NAME = api-service-nuxa-tech
PORT = 3000

.PHONY: run build tidy fmt lint test migrate up down

run:
	go run cmd/main.go	

build:
	go build -o $(APP_NAME) cmd/main.go

tidy:
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...

migrate:
	docker-compose exec app migrate -path db/migrations -database "postgres://postgres:password@db:5432/postgres?sslmode=disable" up

up:
	docker-compose up -d --build

down:
	docker-compose down
