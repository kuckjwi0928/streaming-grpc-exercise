package service

import (
	"context"
	"streaming-grpc-exercise/pkg/model"
	"streaming-grpc-exercise/pkg/repository"
)

type BookService interface {
	GetAllBooks(ctx context.Context) []model.Book
}

type bookServiceImpl struct {
	bookRepo repository.BookRepository
}

func (b *bookServiceImpl) GetAllBooks(ctx context.Context) []model.Book {
	return b.bookRepo.GetAllBooks(ctx)
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookServiceImpl{
		bookRepo: bookRepo,
	}
}
