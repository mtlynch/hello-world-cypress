FROM golang:1.12.4-alpine3.9

RUN mkdir /app
COPY main.go /app/main.go
WORKDIR /app

ENTRYPOINT go run main.go