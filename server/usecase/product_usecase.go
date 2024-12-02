package usecase

import (
	"math"
	"server/model"
	"server/model/dto"
	"server/model/dto/response"
	"server/repository"
)

type ProductUsecase interface {
	CreateProduct(newProduct model.Product) (model.Product, error)
	GetProduct(order string, sort string, page int, limit int) ([]interface{}, dto.Paging, error)
	UpdateProductById(updatedProduct model.Product) (response.UpdatedProductResponse, error)
	DeleteProductById(id string) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func (uc *productUsecase) CreateProduct(newProduct model.Product) (model.Product, error) {
	return uc.repo.CreateProduct(newProduct)
}

func (uc *productUsecase) GetProduct(order string, sort string, page int, limit int) ([]interface{}, dto.Paging, error){
	offset := (page - 1) * limit

	data, totalRows, err := uc.repo.GetProduct(order, sort, limit, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging {
		Page: page,
		TotalPages: int(math.Ceil(float64(totalRows)/float64(limit))),
		TotalRows: totalRows,
		RowsPerPage: limit,
	}

	return data, paging, nil
}

func (uc *productUsecase) UpdateProductById(updatedProduct model.Product) (response.UpdatedProductResponse, error){
	data, err := uc.repo.UpdateProductById(updatedProduct)

	if err != nil {
		return response.UpdatedProductResponse{}, err
	}

	res := response.UpdatedProductResponse {
		Id: data.Id,
		Price: data.Price,
		Name: data.Name,
	}

	return res, nil
}

func (uc *productUsecase) DeleteProductById(id string) error{
	return uc.repo.DeleteProductById(id)
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}