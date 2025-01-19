package repository

import (
	"database/sql"
	"fmt"
	"server/model"
	"server/model/dto/response"
	"server/utils"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(NewUser model.User) (response.UserResponse, error)
	GetAllUser(order string, sort string, limit int, offset int) ([]any, int, error)
	DeleteUserById(id string) (string, error)
}

type userRepository struct {
	db *sql.DB
}

func (pr *userRepository) CreateUser(NewUser model.User) (response.UserResponse, error) {

	NewUser.Id = uuid.NewString()
	now := time.Now().UTC()
	_, err := pr.db.Exec(utils.INSERT_USER,
		NewUser.Id,
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
		Id:        NewUser.Id,
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
