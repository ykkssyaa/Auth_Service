package service

import (
	"AuthService/internal/model"
	"AuthService/internal/repository"
	"AuthService/pkg/logger"
)

type AuthService interface {
	GenerateTokens(guid string) (model.Tokens, error)
}

type AuthServiceImpl struct {
	repo   repository.AuthRepository
	logger *logger.Logger
}

func (a AuthServiceImpl) GenerateTokens(guid string) (model.Tokens, error) {

	refreshToken, err := model.GenerateToken(guid, model.RefreshTokenTTL)

	if err != nil {
		return model.Tokens{}, err
	}

	accessToken, err := model.GenerateToken(guid, model.AccessTokenTTL)

	if err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}
