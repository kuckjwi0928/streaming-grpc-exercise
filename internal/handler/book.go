package handler

import (
	"context"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"streaming-grpc-exercise/api/pb"
	"streaming-grpc-exercise/pkg/model"
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
	lo.ForEach[model.Book](books, func(el model.Book, _ int) {
		_ = stream.Send(&pb.Book{
			Id:     el.Id,
			Author: el.Author,
			Name:   el.Name,
		})
	})
	return nil
}
