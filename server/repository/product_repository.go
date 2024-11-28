package repository

import (
	"database/sql"
)

type ProductRepository interface{

}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}