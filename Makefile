.PHONY: generate
generate:
	protoc --include_imports --include_source_info --proto_path=./googleapis --proto_path=./protos --go_out=./api/pb --go_opt=paths=source_relative \
		--go-grpc_out=./api/pb --go-grpc_opt=paths=source_relative --descriptor_set_out=deployments/k8s/proto.pb `find ./protos -iname "*.proto"`
