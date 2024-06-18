package storage

import "go.mongodb.org/mongo-driver/mongo"

type MongoStorage struct {
	userColl *mongo.Collection
}

func NewMongoStorage(database *mongo.Database) *MongoStorage {
	return &MongoStorage{
		userColl: database.Collection("users"),
	}
}
