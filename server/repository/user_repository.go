package repository

import (
	"database/sql"
	"fmt"
	"server/model"
	"server/model/dto/response"
	"server/utils"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)

	CreateUser(newUser model.User) error
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

func (pr *userRepository) CreateUser(newUser model.User) error {

	_, err := pr.db.Exec(utils.INSERT_USER,
		newUser.Id,
		newUser.Email,
		newUser.Nickname,
		newUser.Password,
		newUser.CreatedAt,
		newUser.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
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
