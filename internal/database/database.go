package database

import (
	"github.com/hashicorp/go-memdb"
	"log"
	"streaming-grpc-exercise/pkg/model"
)

func NewDatabase() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"book": {
				Name: "book",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"author": {
						Name:    "author",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Author"},
					},
				},
			},
		},
	}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	tx := db.Txn(true)
	fixtures := []model.Book{
		{
			Id:     1,
			Name:   "Kotlin in Action",
			Author: "Dmitry Jemerov and Svetlana Isakova",
		},
		{
			Id:     2,
			Name:   "Christian E. Posta",
			Author: "2",
		},
		{
			Id:     3,
			Name:   "Mastering Go",
			Author: "Mihalis Tsoukalos",
		},
	}

	for _, b := range fixtures {
		err := tx.Insert("book", b)
		if err != nil {
			panic(err)
		}
	}

	tx.Commit()

	log.Println("Initial database...")

	if err != nil {
		panic(err)
	}

	return db
}
