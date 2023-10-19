FROM golang:1.20-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o broker-service ./cmd/api

RUN chmod +x ./broker-service

#build tiny image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker-service /app

CMD ["/app/broker-service"]