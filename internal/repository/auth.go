package repository

import (
	"AuthService/internal/model"
	"AuthService/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthRepository interface {
	SaveToken(id, tokenHash string) error
	FindTokens(id string) ([]model.MongoDoc, error)
	DeleteToken(tokenId primitive.ObjectID) error
}

type AuthRepositoryImpl struct {
	mongo  *mongo.Client
	logger *logger.Logger
}

func (r AuthRepositoryImpl) SaveToken(id, tokenHash string) error {

	doc := bson.D{{"user_id", id}, {"token", tokenHash}, {"created_time", time.Now()}}

	_, err := r.mongo.Database(DatabaseName).Collection(CollectionName).InsertOne(context.TODO(), doc)

	return err
}

func (r AuthRepositoryImpl) FindTokens(id string) ([]model.MongoDoc, error) {

	filter := bson.D{{"user_id", id}}

	cursor, err := r.mongo.Database(DatabaseName).Collection(CollectionName).Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var docs []model.MongoDoc
	if err := cursor.All(context.TODO(), &docs); err != nil {
		return nil, err
	}

	return docs, nil
}

func (r AuthRepositoryImpl) DeleteToken(tokenId primitive.ObjectID) error {

	filter := bson.D{{"_id", tokenId}}

	_, err := r.mongo.Database(DatabaseName).Collection(CollectionName).DeleteOne(context.TODO(), filter)

	return err
}
