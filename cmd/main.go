package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"streaming-grpc-exercise/configs"
	"streaming-grpc-exercise/internal/database"
	"streaming-grpc-exercise/internal/handler"
	"streaming-grpc-exercise/pkg/repository"
	"streaming-grpc-exercise/pkg/server"
	"streaming-grpc-exercise/pkg/service"
	"sync"
	"syscall"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.Server.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := server.NewServer()
	repo := repository.NewRepository(database.NewDatabase())
	container := service.NewContainer(repo)

	handler.RegisterBookService(grpcServer, container)
	handler.RegisterHealthCheck(grpcServer)

	wg := sync.WaitGroup{}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		<-signalCh
		grpcServer.GracefulStop()
		wg.Done()
		log.Println("Shutdown server...")
	}()

	log.Println("Starting server...")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	wg.Wait()
}
