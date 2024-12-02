package repository

import (
	"database/sql"
	"fmt"
	"server/model"
	"server/utils"

	"github.com/google/uuid"
)

type ProductRepository interface{
	CreateProduct(NewProduct model.Product) (model.Product, error)
	GetProduct(order string, sort string, limit int, offset int) ([]interface{}, int, error)
	UpdateProductById(updatedProduct model.Product) (model.Product, error)
	DeleteProductById(id string) error
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

func (cr *productRepository) GetProduct(order string, sort string, limit int, offset int) ([]interface{}, int, error){

	query := fmt.Sprintf(utils.SELECT_PRODUCT_WITH_PAGING, order, sort)

	rows, err := cr.db.Query(query, limit, offset)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var products []interface{}
	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.Id,			
			&product.Price,
			&product.Name,
			&product.CreatedAt,
		)

		if err != nil {
			return nil, -1, err
		}

		products = append(products, product)
	}

	var totalRows int
	err = cr.db.QueryRow(utils.SELECT_COUNT_PRODUCT).Scan(&totalRows)
	if err != nil {
		return nil, -1, err
	}

	return products, totalRows, nil
}

func (cr *productRepository) UpdateProductById(updatedProduct model.Product) (model.Product, error){
	_, err := cr.db.Exec(utils.UPDATE_PRODUCT_BY_ID, updatedProduct.Id, updatedProduct.Price, updatedProduct.Name)

	if err != nil {
		return model.Product{}, fmt.Errorf("invalid request body")
	}

	return updatedProduct, nil
}

func (cr *productRepository) DeleteProductById(id string) error{
	_, err := cr.db.Query(utils.DELETE_PRODUCT_BY_ID, id)
	if err != nil {
		return err
	}

	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}