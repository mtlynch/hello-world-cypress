FROM golang:1.12.4-alpine3.9 as hw-builder
LABEL stage=hw-builder
WORKDIR /app
COPY main.go /app/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM scratch
COPY --from=hw-builder /app/main .
ENTRYPOINT [ "/main" ]