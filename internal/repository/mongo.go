package repository

import (
	"AuthService/internal/consts"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DatabaseName = "Auth"
const CollectionName = "Refresh"

func NewMongoClient(dbname, host, port, user, password string) (*mongo.Client, error) {

	uri := fmt.Sprintf("%s://%s:%s@%s:%s/", dbname, user, password, host, port)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client, nil
}

func CreateCollections(client *mongo.Client) error {

	err := client.Database(DatabaseName).CreateCollection(context.TODO(), CollectionName)
	if err != nil {
		return err
	}

	var sec = consts.ExpireAfterSeconds
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"created_time", 1},
		},
		Options: &options.IndexOptions{
			ExpireAfterSeconds: &sec, // 604800 sec = 7 days * 24h * 60 min * 60sec
		},
	}

	_, err = client.Database(DatabaseName).Collection(CollectionName).Indexes().CreateOne(context.TODO(), indexModel)

	return err
}
