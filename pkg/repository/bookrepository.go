package repository

import (
	"context"
	"github.com/hashicorp/go-memdb"
	"streaming-grpc-exercise/pkg/model"
)

type BookRepository interface {
	GetAllBooks(ctx context.Context) []model.Book
}

type bookRepositoryImpl struct {
	db *memdb.MemDB
}

func NewBookRepository(db *memdb.MemDB) BookRepository {
	return bookRepositoryImpl{
		db: db,
	}
}

func (b bookRepositoryImpl) GetAllBooks(ctx context.Context) []model.Book {
	tx := b.db.Txn(false)
	it, _ := tx.Get("book", "id")
	books := make([]model.Book, 0)
	for n := it.Next(); n != nil; n = it.Next() {
		books = append(books, n.(model.Book))
	}
	return books
}
