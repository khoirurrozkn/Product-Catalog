package usecase

import (
	"math"
	"server/model"
	"server/model/dto"
	"server/repository"
)

type FavoriteUsecase interface {
	CreateFavorite(newFavorite model.Favorite) (model.Favorite, error)
	GetAllFavorite(order string, sort string, page int, limit int) ([]any, dto.Paging, error)
	DeleteFavoriteById(id string) (string, error)
}

type favoriteUsecase struct {
	repo repository.FavoriteRepository
}

func (pu *favoriteUsecase) CreateFavorite(newFavorite model.Favorite) (model.Favorite, error) {
	return pu.repo.CreateFavorite(newFavorite)
}

func (pu *favoriteUsecase) GetAllFavorite(order string, sort string, page int, limit int) ([]any, dto.Paging, error) {
	offset := (page - 1) * limit

	data, totalRows, err := pu.repo.GetAllFavorite(order, sort, limit, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging{
		Page:        page,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
		TotalRows:   totalRows,
		RowsPerPage: limit,
	}

	return data, paging, nil
}

func (pu *favoriteUsecase) DeleteFavoriteById(id string) (string, error) {
	return pu.repo.DeleteFavoriteById(id)
}

func NewFavoriteUsecase(repo repository.FavoriteRepository) FavoriteUsecase {
	return &favoriteUsecase{
		repo: repo,
	}
}
