package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoDoc struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserID  string             `bson:"user_id"`
	Token   string             `bson:"token"`
	Created time.Time          `bson:"created_time"`
}
