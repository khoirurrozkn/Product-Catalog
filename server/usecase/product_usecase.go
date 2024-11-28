package usecase

import (
	"server/model"
	"server/repository"
)

type ProductUsecase interface {
	CreateProduct(newCustomer model.Product) (model.Product, error)
}

type productUsecase struct {
	repo repository.ProductRepository
}

func (uc *productUsecase) CreateProduct(newProduct model.Product) (model.Product, error) {
	return uc.repo.CreateProduct(newProduct)
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}