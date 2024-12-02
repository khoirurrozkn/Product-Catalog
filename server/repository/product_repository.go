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
	GetProduct(order string, sort string, limit int, offset int) ([]any, int, error)
	UpdateProductById(updatedProduct model.Product) (model.Product, error)
	DeleteProductById(id string) (string, error)
}

type productRepository struct {
	db *sql.DB
}

func (pr *productRepository) CreateProduct(NewProduct model.Product) (model.Product, error) {
	NewProduct.Id = uuid.NewString()
	
	data := model.Product{}
	err := pr.db.QueryRow(utils.INSERT_PRODUCT,
		NewProduct.Id,
		NewProduct.ImgUrl,
		NewProduct.Price,
		NewProduct.Name,
	).Scan(&data.Id, &data.ImgUrl, &data.Price, &data.Name, &data.CreatedAt)

	if err != nil {
		return model.Product{}, err
	}

	return data, err
}

func (pr *productRepository) GetProduct(order string, sort string, limit int, offset int) ([]any, int, error){

	query := fmt.Sprintf(utils.SELECT_PRODUCT_WITH_PAGING, order, sort)

	rows, err := pr.db.Query(query, limit, offset)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var products []any
	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.Id,		
			&product.ImgUrl,	
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
	err = pr.db.QueryRow(utils.SELECT_COUNT_PRODUCT).Scan(&totalRows)
	if err != nil {
		return nil, -1, err
	}

	return products, totalRows, nil
}

func (pr *productRepository) UpdateProductById(updatedProduct model.Product) (model.Product, error){
	_, err := pr.db.Exec(utils.UPDATE_PRODUCT_BY_ID, updatedProduct.Id, updatedProduct.ImgUrl, updatedProduct.Price, updatedProduct.Name)

	if err != nil {
		return model.Product{}, err
	}

	return updatedProduct, nil
}

func (pr *productRepository) DeleteProductById(id string) (string, error){
	_, err := pr.db.Query(utils.DELETE_PRODUCT_BY_ID, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}