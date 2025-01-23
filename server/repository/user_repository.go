package repository

import (
	"database/sql"
	"fmt"
	"server/model"
	"server/model/dto/request"
	"server/model/dto/response"
	"server/utils"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	GetUserByNickname(nickname string) (model.User, error)

	CreateUser(NewUser request.UserRegister) (response.UserResponse, error)
	GetAllUser(order string, sort string, limit int, offset int) ([]any, int, error)
	DeleteUserById(id string) (string, error)
}

type userRepository struct {
	db *sql.DB
}

func (pr *userRepository) GetUserByEmail(email string) (model.User, error) {

	var user model.User
	err := pr.db.QueryRow(utils.SELECT_USER_BY_EMAIL, email).Scan(
		&user.Id,
		&user.Email,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (pr *userRepository) GetUserByNickname(nickname string) (model.User, error) {

	var user model.User
	err := pr.db.QueryRow(utils.SELECT_USER_BY_NICKNAME, nickname).Scan(
		&user.Id,
		&user.Email,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (pr *userRepository) CreateUser(NewUser request.UserRegister) (response.UserResponse, error) {

	id := uuid.NewString()
	now := time.Now().UTC()
	_, err := pr.db.Exec(utils.INSERT_USER,
		id,
		NewUser.Email,
		NewUser.Nickname,
		NewUser.Password,
		now,
		now,
	)
	if err != nil {
		return response.UserResponse{}, err
	}
	user := response.UserResponse{
		Id:        id,
		Email:     NewUser.Email,
		Nickname:  NewUser.Nickname,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return user, err
}

func (pr *userRepository) GetAllUser(order string, sort string, limit int, offset int) ([]any, int, error) {

	query := fmt.Sprintf(utils.SELECT_USER_WITH_PAGING, order, sort)

	rows, err := pr.db.Query(query, limit, offset)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var users []any
	for rows.Next() {
		var user response.UserResponse
		err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.Nickname,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, -1, err
		}

		users = append(users, user)
	}

	var totalRows int
	err = pr.db.QueryRow(utils.SELECT_COUNT_USER).Scan(&totalRows)
	if err != nil {
		return nil, -1, err
	}

	return users, totalRows, nil
}

func (pr *userRepository) DeleteUserById(id string) (string, error) {
	_, err := pr.db.Exec(utils.DELETE_USER_BY_ID, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
