FROM golang:1.17-alpine3.14 AS build
WORKDIR /app

COPY cmd /app/cmd
COPY entities /app/entities
COPY infrastructure /app/infrastructure
COPY user_case /app/user_case
COPY .env /app/.env
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
COPY main.go /app/main.go

RUN go mod download
# RUN go build -o build ./src/main.go
ENTRYPOINT go run main.go httpserver