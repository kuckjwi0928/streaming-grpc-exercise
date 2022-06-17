# GRPC Streaming exercise
## Clone
```shell
git clone --recursive https://github.com/kuckjwi0928/streaming-grpc-exercise
```

## Debug on kubernetes (skaffold)
```shell
skaffold dev
```

## grpcurl call example
```shell
grpcurl -plaintext localhost:8080 BookService/ListBook
```
