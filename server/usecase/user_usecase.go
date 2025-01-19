package usecase

import (
	"math"
	"server/model"
	"server/model/dto"
	"server/model/dto/response"
	"server/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(newUser model.User) (response.UserResponse, error)
	GetAllUser(order string, sort string, page int, limit int) ([]any, dto.Paging, error)
	DeleteUserById(id string) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func (pu *userUsecase) CreateUser(newUser model.User) (response.UserResponse, error) {
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return response.UserResponse{}, err
	}

	newUser.Password = string(hashedPassword)

	return pu.repo.CreateUser(newUser)
}

func (pu *userUsecase) GetAllUser(order string, sort string, page int, limit int) ([]any, dto.Paging, error) {
	offset := (page - 1) * limit

	data, totalRows, err := pu.repo.GetAllUser(order, sort, limit, offset)
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

func (pu *userUsecase) DeleteUserById(id string) (string, error) {
	return pu.repo.DeleteUserById(id)
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}
