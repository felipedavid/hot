package storage

import (
	"context"
	"errors"

	"github.com/felipedavid/hot/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRoom(ctx context.Context, id string) (*types.Room, error) {
	var room types.Room

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = roomsColl.FindOne(ctx, bson.M{"_id": oid}).Decode(&room)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func GetRooms(ctx context.Context) ([]*types.Room, error) {
	cur, err := roomsColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var rooms []*types.Room
	err = cur.All(ctx, &rooms)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func InsertRoom(ctx context.Context, room *types.Room) error {
	res, err := roomsColl.InsertOne(ctx, room)
	if err != nil {
		return err
	}

	uid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("unable to cast objectid")
	}
	room.ID = uid.Hex()

	return nil
}

func DeleteRoom(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := roomsColl.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func UpdateRoom(ctx context.Context, room *types.Room) error {
	update := bson.M{
		"$set": bson.M{
			"type":       room.Type,
			"base_price": room.BasePrice,
			"price":      room.Price,
			"hotel_id":   room.HotelID,
		},
	}

	oid, err := primitive.ObjectIDFromHex(room.ID)
	if err != nil {
		return err
	}

	_, err = roomsColl.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return err
	}

	return nil
}
