include .env
export

install-deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest  
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/vektra/mockery/v2@v2.35.4

get-deps:
	go get -u github.com/go-faker/faker/v4
	go get -u github.com/joho/godotenv
	go get -u github.com/google/uuid
	go get -u github.com/lib/pq
	go get -u github.com/golang-jwt/jwt/v5
	go get -u github.com/swaggo/http-swagger

swag:
	swag init -g internal/app/app.go

up-goose-migration:
	cd postgresql/schema && goose postgres $(DB_URL) up && cd ../../

generate-sqlc:
	mkdir -p internal/repository/sqlc/generated
	sqlc generate

down-goose-migration:
	cd postgresql/schema && goose postgres $(DB_URL) down && cd ../../

build:
	docker build .

run:
	docker compose up

