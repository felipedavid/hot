package storage

import "go.mongodb.org/mongo-driver/mongo"

var userColl *mongo.Collection

func Init(database *mongo.Database) {
	userColl = database.Collection("users")
}
