package usecase

import (
	"server/model"
	"server/repository"
)

type ProductUsecase interface {
	CreateProduct(newProduct model.Product) (model.Product, error)
	GetProduct() ([]interface{}, error)
	UpdateProductById(updatedProduct model.Product) (model.Product, error)
	DeleteProductById(id string) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func (uc *productUsecase) CreateProduct(newProduct model.Product) (model.Product, error) {
	return uc.repo.CreateProduct(newProduct)
}

func (uc *productUsecase) GetProduct() ([]interface{}, error){
	return uc.repo.GetProduct()
}

func (uc *productUsecase) UpdateProductById(updatedProduct model.Product) (model.Product, error){
	return uc.repo.UpdateProductById(updatedProduct)
}

func (uc *productUsecase) DeleteProductById(id string) error{
	return uc.repo.DeleteProductById(id)
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}