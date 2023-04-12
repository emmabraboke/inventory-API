package tokenService

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type token struct {
	Email  string
	Id     string
	jwt.StandardClaims
}

type TokenService interface {
	CreateToken(id,email string) (string, string, error)
	ValidateToken(token string) (*token, error)
}

type tokenSrv struct {
	SecretKey string
}

func (t *tokenSrv) CreateToken(id, email string) (string, string, error) {
	tokenDetails := &token{
		Email:  email,
		Id:     id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshTokenDetails := &token{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(60)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenDetails).SignedString([]byte(t.SecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenDetails).SignedString([]byte(t.SecretKey))

	if err != nil {
		log.Println(err)
		return "", "", err
	}

	return token, refreshToken, err
}

func (t *tokenSrv) ValidateToken(tokenUrl string) (*token, error) {
	data, err := jwt.ParseWithClaims(
		tokenUrl,
		&token{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(t.SecretKey), nil
		},
	)

	claims, ok := data.Claims.(*token)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, err
	}

	return claims, err

}

func NewTokenSrv(secret string) TokenService {
	return &tokenSrv{secret}
}
