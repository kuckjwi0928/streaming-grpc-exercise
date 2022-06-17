package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer() *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)
	return server
}
