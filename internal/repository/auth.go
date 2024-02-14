package repository

import (
	"AuthService/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
}

type AuthRepositoryImpl struct {
	mongo  *mongo.Client
	logger *logger.Logger
}
