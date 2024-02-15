package service

import (
	"AuthService/internal/repository"
	"AuthService/pkg/logger"
)

type Services struct {
	AuthService AuthService
}

func NewServices(repository *repository.Repository, logger *logger.Logger) *Services {
	return &Services{
		AuthService: AuthServiceImpl{logger: logger, repo: repository.AuthRepository},
	}
}
