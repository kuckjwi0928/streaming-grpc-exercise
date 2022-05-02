package service

import "streaming-grpc-exercise/pkg/repository"

type Container struct {
	repo *repository.Repository

	BookService BookService
}

func NewContainer(repo *repository.Repository) *Container {
	return &Container{
		BookService: NewBookService(repo.Book),
	}
}
