package usecase

import (
	"errors"
	"math"
	"server/model"
	"server/model/dto"
	"server/model/dto/request"
	"server/repository"
	"time"

	"github.com/google/uuid"
)

type FavoriteUsecase interface {
	CreateFavorite(newFavorite request.Favorite, userId string) (model.Favorite, error)
	GetAllFavorite(userId string, order string, sort string, page int, limit int) ([]any, dto.Paging, error)
	DeleteFavoriteById(userId string, favoriteId string) (string, error)
}

type favoriteUsecase struct {
	repo repository.FavoriteRepository
}

func (pu *favoriteUsecase) CreateFavorite(newFavorite request.Favorite, UserId string) (model.Favorite, error) {

	data := model.Favorite{
		Id: uuid.NewString(),
		UserId: UserId,
		ProductId: newFavorite.ProductId,
		CreatedAt: time.Now().UTC(),
	}

	err := pu.repo.CreateFavorite(data)
	if err != nil {
		return model.Favorite{}, err
	}

	return data, nil
}

func (pu *favoriteUsecase) GetAllFavorite(userId string, order string, sort string, page int, limit int) ([]any, dto.Paging, error) {
	offset := (page - 1) * limit

	data, totalRows, err := pu.repo.GetAllFavorite(userId, order, sort, limit, offset)
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

func (pu *favoriteUsecase) DeleteFavoriteById(userId string, favoriteId string) (string, error) {
	findById, err := pu.repo.GetById(favoriteId)
	if err != nil {
		return "", err
	}

	if findById.UserId != userId {
		return "", errors.New("hayooo mau ngapain")
	}

	return pu.repo.DeleteFavoriteById(favoriteId)
}

func NewFavoriteUsecase(repo repository.FavoriteRepository) FavoriteUsecase {
	return &favoriteUsecase{
		repo: repo,
	}
}
