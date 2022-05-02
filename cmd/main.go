package main

import (
	"fmt"
	"log"
	"net"
	"streaming-grpc-exercise/internal/handler"
	"streaming-grpc-exercise/pkg/database"
	"streaming-grpc-exercise/pkg/repository"
	"streaming-grpc-exercise/pkg/server"
	"streaming-grpc-exercise/pkg/service"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		panic(err)
	}

	grpcServer := server.NewServer()
	repo := repository.NewRepository(database.NewDatabase())
	container := service.NewContainer(repo)

	handler.RegisterBookService(grpcServer, container)

	log.Println("Starting server...")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
