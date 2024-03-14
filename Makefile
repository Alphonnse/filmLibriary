include .env

install-deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest  
	go install github.com/pressly/goose/v3/cmd/goose@latest

get-deps:
	go get -u github.com/joho/godotenv
	go get -u github.com/google/uuid
	go get -u github.com/lib/pq
	go get -u github.com/golang-jwt/jwt/v5

generate-repository:
	make generate-sqlc && make up-goose-migration

generate-sqlc:
	mkdir -p internal/repository/sqlc/generated
	sqlc generate

up-goose-migration:
	cd postgresql/schema && goose postgres postgres://alphonnnse:@localhost:5432/FilmLibriary up && cd ../../

down-goose-migration:
	cd postgresql/schema && goose postgres postgres://alphonnnse:@localhost:5432/FilmLibriary down && cd ../../
