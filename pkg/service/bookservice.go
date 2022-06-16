package service

import (
	"context"
	"github.com/samber/lo"
	"streaming-grpc-exercise/pkg/model"
	"streaming-grpc-exercise/pkg/repository"
)

type BookService interface {
	GetAllBooks(ctx context.Context) []model.Book
	GetBook(ctx context.Context, id int64) *model.Book
}

type bookServiceImpl struct {
	bookRepo repository.BookRepository
}

func (b *bookServiceImpl) GetAllBooks(ctx context.Context) []model.Book {
	return b.bookRepo.GetAllBooks(ctx)
}

func (b *bookServiceImpl) GetBook(ctx context.Context, id int64) *model.Book {
	book, _ := lo.Find[model.Book](b.bookRepo.GetAllBooks(ctx), func(book model.Book) bool {
		return book.Id == id
	})
	return &book
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookServiceImpl{
		bookRepo: bookRepo,
	}
}
