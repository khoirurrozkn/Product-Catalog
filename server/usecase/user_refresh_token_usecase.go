package usecase

import (
	"server/repository"
	"server/utils/common"
	"time"

	"github.com/google/uuid"
)

type UserRefreshTokenUsecase interface {
	CreateUserRefreshToken(userId string, role string, email string, tokenType string) (string, error)
	GetUserRefreshTokenById(token string) (any, error)
}

type userRefreshTokenUsecase struct {
	repo repository.UserRefreshTokenRepository
	jwtToken common.JwtToken
}

func (urtu *userRefreshTokenUsecase) CreateUserRefreshToken(userId string, role string, email string, tokenType string) (string, error) {
	token, lifeTime, err := urtu.jwtToken.GenerateTokenJwt(userId, role, email, "refreshToken")
	if err != nil {
		return "", err
	}

	tokenId := uuid.NewString()
	now := time.Now().UTC()
	oneWeekLater := now.Add(lifeTime * time.Minute)

	err = urtu.repo.CreateUserRefreshToken(tokenId, userId, token, oneWeekLater)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (urtu *userRefreshTokenUsecase) GetUserRefreshTokenById(token string) (any, error) {
	return urtu.repo.GetUserRefreshTokenByToken(token)
}

func NewUserRefreshTokenUsecase(repo repository.UserRefreshTokenRepository, jwtToken common.JwtToken) UserRefreshTokenUsecase {
	return &userRefreshTokenUsecase{
		repo: repo,
		jwtToken: jwtToken,
	}
}
