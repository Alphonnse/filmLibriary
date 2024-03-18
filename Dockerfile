FROM golang:1.21.1

WORKDIR /app

RUN apt-get update

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make install-deps
RUN make get-deps

RUN go generate ./...

COPY .env .
ENV SERVER_HOST=${SERVER_HOST} \
    SERVER_PORT=${SERVER_PORT} \
    DB_URL=${DB_URL} \
    SECRET=${SECRET} 

RUN go build -o app ./cmd/filmLibriary/main.go

EXPOSE 8000

CMD ["./app"]
