FROM golang:1.18.1-alpine3.15 AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o streaming-grpc-exercise cmd/main.go

FROM alpine:3.15.4
WORKDIR /app
COPY --from=builder /app/streaming-grpc-exercise ./
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
CMD ["sh", "-c", "/app/streaming-grpc-exercise"]
