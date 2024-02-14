package repository

import (
	"AuthService/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	AuthRepository AuthRepository
}

func NewRepository(mongo *mongo.Client, logger *logger.Logger) *Repository {
	return &Repository{
		AuthRepository: AuthRepositoryImpl{logger: logger, mongo: mongo},
	}
}
