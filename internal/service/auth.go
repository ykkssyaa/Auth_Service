package service

import (
	"AuthService/internal/model"
	"AuthService/internal/repository"
	"AuthService/pkg/logger"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	GenerateTokens(guid string) (model.Tokens, error)
}

type AuthServiceImpl struct {
	repo   repository.AuthRepository
	logger *logger.Logger
}

func (a AuthServiceImpl) GenerateTokens(guid string) (model.Tokens, error) {

	accessToken, err := model.GenerateToken(guid, model.AccessTokenTTL)

	if err != nil {
		return model.Tokens{}, err
	}

	refreshToken := generateRefreshHash(guid)

	hash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), 10)

	if err != nil {
		return model.Tokens{}, err
	}

	err = a.repo.SaveToken(guid, string(hash))

	return model.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, err
}

func generateRefreshHash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))

	time.Sleep(1 * time.Millisecond)
	hash.Write([]byte(time.Now().String()))

	token := fmt.Sprintf("%x", hash.Sum(nil))

	return token[len(token)-30:]
}

func (a AuthServiceImpl) RefreshTokens(refreshToken, guid string) (model.Tokens, error) {
	return model.Tokens{}, nil
}
