FROM golang:1.18.1-alpine3.15 AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o streaming-grpc-exercise cmd/main.go

FROM alpine:3.15.4
WORKDIR /app
COPY --from=builder /app/streaming-grpc-exercise ./
EXPOSE 8080
