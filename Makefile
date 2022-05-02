.PHONY: generate
generate:
	protoc --proto_path=./protos --go_out=./api/pb --go_opt=paths=source_relative \
		--go-grpc_out=./api/pb --go-grpc_opt=paths=source_relative `find ./protos -iname "*.proto"`
