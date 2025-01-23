package usecase

import (
	"errors"
	"math"
	"server/model/dto"
	"server/model/dto/request"
	"server/model/dto/response"
	"server/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	LoginWithEmail(user request.UserLogin) (response.UserResponse, error)
	LoginWithNickname(user request.UserLogin) (response.UserResponse, error)

	CreateUser(newUser request.UserRegister) (response.UserResponse, error)
	GetAllUser(order string, sort string, page int, limit int) ([]any, dto.Paging, error)
	DeleteUserById(id string) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func (uu *userUsecase) LoginWithEmail(user request.UserLogin) (response.UserResponse, error){

	findUser, err := uu.repo.GetUserByEmail(user.EmailOrNickname)
	if err != nil {
		return response.UserResponse{}, err
	}

	if valid := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password)); valid != nil {
		return response.UserResponse{}, errors.New("incorrect password")
	}

	data := response.UserResponse{
		Id:        findUser.Id,
		Email:     findUser.Email,
		Nickname:  findUser.Nickname,
		CreatedAt: findUser.CreatedAt,
		UpdatedAt: findUser.UpdatedAt,
	}
	
	return data, nil
}

func (uu *userUsecase) LoginWithNickname(user request.UserLogin) (response.UserResponse, error){

	findUser, err := uu.repo.GetUserByNickname(user.EmailOrNickname)
	if err != nil {
		return response.UserResponse{}, err
	}

	if valid := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password)); valid != nil {
		return response.UserResponse{}, errors.New("incorrect password")
	}

	data := response.UserResponse{
		Id:        findUser.Id,
		Email:     findUser.Email,
		Nickname:  findUser.Nickname,
		CreatedAt: findUser.CreatedAt,
		UpdatedAt: findUser.UpdatedAt,
	}
	
	return data, nil
}

func (uu *userUsecase) CreateUser(newUser request.UserRegister) (response.UserResponse, error) {
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return response.UserResponse{}, err
	}

	newUser.Password = string(hashedPassword)

	return uu.repo.CreateUser(newUser)
}

func (uu *userUsecase) GetAllUser(order string, sort string, page int, limit int) ([]any, dto.Paging, error) {
	offset := (page - 1) * limit

	data, totalRows, err := uu.repo.GetAllUser(order, sort, limit, offset)
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

func (uu *userUsecase) DeleteUserById(id string) (string, error) {
	return uu.repo.DeleteUserById(id)
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}
