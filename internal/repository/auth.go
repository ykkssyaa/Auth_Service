package repository

import (
	"AuthService/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthRepository interface {
	SaveToken(id, tokenHash string) error
}

type AuthRepositoryImpl struct {
	mongo  *mongo.Client
	logger *logger.Logger
}

func (r AuthRepositoryImpl) SaveToken(id, tokenHash string) error {

	doc := bson.D{{"id", id}, {"token", tokenHash}, {"created_time", time.Now()}}

	_, err := r.mongo.Database(DatabaseName).Collection(CollectionName).InsertOne(context.TODO(), doc)

	return err
}
