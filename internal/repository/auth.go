package repository

import (
	"AuthService/internal/consts"
	"AuthService/internal/model"
	se "AuthService/pkg/error"
	"AuthService/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
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

	if _, err := r.mongo.Database(DatabaseName).Collection(CollectionName).InsertOne(context.TODO(), doc); err != nil {
		return &se.ResponseError{Message: consts.ErrSaveToken, Code: http.StatusInternalServerError}
	}
	return nil
}

func (r AuthRepositoryImpl) FindTokens(id string) ([]model.MongoDoc, error) {

	filter := bson.D{{"user_id", id}}

	cursor, err := r.mongo.Database(DatabaseName).Collection(CollectionName).Find(context.TODO(), filter)

	if err != nil {
		return nil, &se.ResponseError{Message: consts.ErrFindToken, Code: http.StatusInternalServerError}
	}

	var docs []model.MongoDoc
	if err := cursor.All(context.TODO(), &docs); err != nil {
		return nil, &se.ResponseError{Message: consts.ErrParseToken, Code: http.StatusInternalServerError}
	}

	return docs, nil
}

func (r AuthRepositoryImpl) DeleteToken(tokenId primitive.ObjectID) error {

	filter := bson.D{{"_id", tokenId}}

	if _, err := r.mongo.Database(DatabaseName).Collection(CollectionName).DeleteOne(context.TODO(), filter); err != nil {
		return &se.ResponseError{Message: consts.ErrDeleteToken, Code: http.StatusInternalServerError}
	}

	return nil
}
