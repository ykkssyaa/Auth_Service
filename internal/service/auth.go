package service

import (
	"AuthService/internal/model"
	"AuthService/internal/repository"
	"AuthService/pkg/logger"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	}, nil
}

func generateRefreshHash(token string) string {
	hash := sha256.New()
	hash.Write([]byte(token))

	return fmt.Sprintf("%x", hash.Sum([]byte(nil)))
}
