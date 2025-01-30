package usecase

import (
	"errors"
	"math"
	"server/model"
	"server/model/dto"
	"server/model/dto/request"
	"server/model/dto/response"
	"server/repository"
	"server/utils/common"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	LoginUser(user request.UserLogin) (response.UserCredential, error)

	CreateUser(newUser request.UserRegister) (response.UserResponse, error)
	GetAllUser(order string, sort string, page int, limit int) ([]any, dto.Paging, error)
	DeleteUserById(id string) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
	userRefreshToken UserRefreshTokenUsecase
	jwtToken common.JwtToken
}

func (uu *userUsecase) LoginUser(user request.UserLogin) (response.UserCredential, error){

	findUser, err := uu.repo.GetUserByEmail(user.Email)
	if err != nil {
		return response.UserCredential{}, err
	}

	if valid := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password)); valid != nil {
		return response.UserCredential{}, errors.New("incorrect password")
	}

	accessToken, _, err := uu.jwtToken.GenerateTokenJwt(findUser.Id, "User", findUser.Email, "accessToken")
	if err != nil {
		return response.UserCredential{}, err
	}

	refreshToken, err := uu.userRefreshToken.CreateUserRefreshToken(findUser.Id, "User", findUser.Email, "refreshToken")
	if err != nil {
		return response.UserCredential{}, err
	}
	
	return response.UserCredential{
		User: response.UserResponse{
			Id:        findUser.Id,
			Email:     findUser.Email,
			Nickname:  findUser.Nickname,
			CreatedAt: findUser.CreatedAt,
			UpdatedAt: findUser.UpdatedAt,
		},
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *userUsecase) CreateUser(newUser request.UserRegister) (response.UserResponse, error) {
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return response.UserResponse{}, err
	}

	now := time.Now().UTC()
	data := model.User {
		Id: uuid.NewString(),
		Email: newUser.Email,
		Nickname: newUser.Nickname,
		Password: string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = uu.repo.CreateUser(data)
	if err != nil {
		return response.UserResponse{}, nil
	}

	return response.UserResponse{
		Id: data.Id,
		Email: data.Email,
		Nickname: data.Nickname,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
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

func NewUserUsecase(repo repository.UserRepository, userRefreshToken UserRefreshTokenUsecase, jwt_token common.JwtToken) UserUsecase {
	return &userUsecase{
		repo: repo,
		userRefreshToken: userRefreshToken,
		jwtToken: jwt_token,
	}
}
