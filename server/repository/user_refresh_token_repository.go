package repository

import (
	"database/sql"
	"server/utils"
	"time"
)

type UserRefreshTokenRepository interface {
	CreateUserRefreshToken(id string, userId string, token string, future time.Time) error
	GetUserRefreshTokenByToken(token string) (string, error)
}

type userRefreshTokenRepository struct {
	db *sql.DB
}

func (pr *userRefreshTokenRepository) CreateUserRefreshToken(id string, userId string, token string, future time.Time) error {

	_, err := pr.db.Exec(utils.INSERT_USER,
		id,
		userId,
		token,
		future,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRefreshTokenRepository) GetUserRefreshTokenByToken(token string) (string, error) {
	_, err := ur.db.Exec(utils.GET_USER_REFRESH_TOKEN_BY_TOKEN, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewUserRefreshTokenRepository(db *sql.DB) UserRefreshTokenRepository {
	return &userRefreshTokenRepository{
		db: db,
	}
}
