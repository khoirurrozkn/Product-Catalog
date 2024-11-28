package usecase

import (
	"server/repository"
)

type ProductUsecase interface {

}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}