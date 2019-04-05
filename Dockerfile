FROM golang:1.12

RUN mkdir /app
COPY main.go /app/main.go
WORKDIR /app

ENTRYPOINT go run main.go