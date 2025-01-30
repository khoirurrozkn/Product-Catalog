package repository

import (
	"database/sql"
	"fmt"
	"server/model"
	"server/utils"
	"time"
)

type FavoriteRepository interface {
	CreateFavorite(id string, userId string, productId string, now time.Time) error
	GetAllFavorite(order string, sort string, limit int, offset int) ([]any, int, error)
	DeleteFavoriteById(id string) (string, error)
}

type favoriteRepository struct {
	db *sql.DB
}

func (pr *favoriteRepository) CreateFavorite(id string, userId string, productId string, now time.Time) error {

	_, err := pr.db.Exec(utils.INSERT_FAVORITE,
		id,
		userId,
		productId,
		now,
	)
	if err != nil {
		return err
	}
	return err
}

func (pr *favoriteRepository) GetAllFavorite(order string, sort string, limit int, offset int) ([]any, int, error) {

	query := fmt.Sprintf(utils.SELECT_FAVORITE_WITH_PAGING, order, sort)

	rows, err := pr.db.Query(query, limit, offset)
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
