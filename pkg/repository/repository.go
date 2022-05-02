package repository

import "github.com/hashicorp/go-memdb"

type Repository struct {
	Book BookRepository
}

func NewRepository(db *memdb.MemDB) *Repository {
	return &Repository{
		Book: NewBookRepository(db),
	}
}
