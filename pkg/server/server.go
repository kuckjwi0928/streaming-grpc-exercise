package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer() *grpc.Server {
	server := grpc.NewServer()
	registerReflection(server)
	return server
}

func registerReflection(server *grpc.Server) {
	reflection.Register(server)
}
