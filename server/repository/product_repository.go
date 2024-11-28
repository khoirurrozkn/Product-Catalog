package repository

import (
	"server/model"
	"server/utils"
	"database/sql"
	"github.com/google/uuid"
)

type ProductRepository interface{
	CreateProduct(NewCustomer model.Product) (model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func (cr *productRepository) CreateProduct(NewProduct model.Product) (model.Product, error) {
	NewProduct.Id = uuid.NewString()
	
	_, err := cr.db.Exec(utils.INSERT_PRODUCT,
		NewProduct.Id,
		NewProduct.Price,
		NewProduct.Name,
	)

	if err != nil {
		return model.Product{}, err
	}

	return NewProduct, err
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}