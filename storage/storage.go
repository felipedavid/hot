package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var usersColl *mongo.Collection
var hotelsColl *mongo.Collection
var roomsColl *mongo.Collection

func Init(database *mongo.Database) {
	usersColl = database.Collection("users")
	hotelsColl = database.Collection("hotels")
	roomsColl = database.Collection("rooms")
}

func Drop(ctx context.Context) error {
	return nil
}
