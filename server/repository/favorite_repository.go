package repository

import (
	"database/sql"
	"fmt"
	"server/model"
	"server/utils"
)

type FavoriteRepository interface {
	GetById(favoriteId string) (model.Favorite, error)

	CreateFavorite(newFavorite model.Favorite) error
	GetAllFavorite(userId string, order string, sort string, limit int, offset int) ([]any, int, error)
	DeleteFavoriteById(id string) (string, error)
}

type favoriteRepository struct {
	db *sql.DB
}

func (pr *favoriteRepository) GetById(favoriteId string) (model.Favorite, error) {

	var data model.Favorite
	err := pr.db.QueryRow(utils.SELECT_FAVORITE_BY_ID, favoriteId).Scan(
		data.Id,
		data.UserId,
		data.ProductId,
		data.CreatedAt,
	)
	if err != nil {
		return model.Favorite{}, err
	}
	return data, nil
}

func (pr *favoriteRepository) CreateFavorite(newFavorite model.Favorite) error {

	_, err := pr.db.Exec(utils.INSERT_FAVORITE,
		newFavorite.Id,
		newFavorite.UserId,
		newFavorite.ProductId,
		newFavorite.CreatedAt,
	)
	if err != nil {
		return err
	}
	return err
}

func (pr *favoriteRepository) GetAllFavorite(userId string, order string, sort string, limit int, offset int) ([]any, int, error) {

	query := fmt.Sprintf(utils.SELECT_FAVORITE_WITH_PAGING, order, sort)

	rows, err := pr.db.Query(query, userId, limit, offset)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var favorites []any
	for rows.Next() {
		var favorite model.Favorite
		err = rows.Scan(
			&favorite.Id,
			&favorite.UserId,
			&favorite.ProductId,
			&favorite.CreatedAt,
		)

		if err != nil {
			return nil, -1, err
		}

		favorites = append(favorites, favorite)
	}

	var totalRows int
	err = pr.db.QueryRow(utils.SELECT_COUNT_FAVORITE).Scan(&totalRows)
	if err != nil {
		return nil, -1, err
	}

	return favorites, totalRows, nil
}

func (pr *favoriteRepository) DeleteFavoriteById(id string) (string, error) {
	_, err := pr.db.Exec(utils.DELETE_FAVORITE_BY_ID, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func NewFavoriteRepository(db *sql.DB) FavoriteRepository {
	return &favoriteRepository{
		db: db,
	}
}
