package service

import (
	"AuthService/internal/repository"
	"AuthService/pkg/logger"
)

type AuthService interface {
}

type AuthServiceImpl struct {
	repo   repository.AuthRepository
	logger *logger.Logger
}
