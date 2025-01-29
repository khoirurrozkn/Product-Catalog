package common

import (
	"errors"
	"server/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaim struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	Email string `json:"email"`
}

type JwtToken interface {
	GenerateTokenJwt(userId string, role string, email string, tokenType string) (string, time.Duration, error)
	VerifyToken(token string) (jwt.MapClaims, error)
}

type jwtToken struct {
	config config.TokenConfig
}

func (cfg *jwtToken) GenerateTokenJwt(userId string, role string, email string, tokenType string) (string, time.Duration, error) {

	var lifeTime time.Duration
	if tokenType == "accessToken" {
		lifeTime = cfg.config.AccessTokenLifeTime
	}else {
		lifeTime = cfg.config.RefreshTokenLifeTime
	}

	claims := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    cfg.config.IssuerName,
			ExpiresAt: time.Now().UTC().Add(lifeTime).Unix(),
		},
		UserId: userId,
		Role:   role,
		Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(cfg.config.JwtSignatureKey)

	if err != nil {
		return "", 0, err
	}
	return signedToken, lifeTime, nil
}

func (cfg *jwtToken) VerifyToken(token_string string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (any, error) {
		return cfg.config.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse map claims or token is not valid")
	}

	if !claims.VerifyIssuer(cfg.config.IssuerName, true) {
		return nil, errors.New("failed to verify issuer name")
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}

func NewJwtToken(token_config config.TokenConfig) JwtToken {
	return &jwtToken{
		config: token_config,
	}
}
