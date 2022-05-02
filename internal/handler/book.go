package handler

import (
	"context"
	"google.golang.org/grpc"
	"streaming-grpc-exercise/api/pb"
	"streaming-grpc-exercise/pkg/service"
	"time"
)

type bookHandler struct {
	container *service.Container
	pb.UnimplementedBookServiceServer
}

func RegisterBookService(server *grpc.Server, container *service.Container) {
	pb.RegisterBookServiceServer(server, &bookHandler{
		container: container,
	})
}

func (b *bookHandler) ListBook(_ *pb.EmptyBookRequest, stream pb.BookService_ListBookServer) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	books := b.container.BookService.GetAllBooks(ctx)
	for _, book := range books {
		err := stream.Send(&pb.Book{
			Id:     book.Id,
			Author: book.Author,
			Name:   book.Name,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
