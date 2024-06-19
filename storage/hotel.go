package storage

import (
	"context"
	"errors"

	"github.com/felipedavid/hot/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetHotel(ctx context.Context, id string) (*types.Hotel, error) {
	var hotel types.Hotel

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = hotelsColl.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel)
	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func GetHotels(ctx context.Context) ([]*types.Hotel, error) {
	cur, err := hotelsColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var hotels []*types.Hotel
	err = cur.All(ctx, &hotels)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func InsertHotel(ctx context.Context, hotel *types.Hotel) error {
	res, err := hotelsColl.InsertOne(ctx, hotel)
	if err != nil {
		return err
	}

	uid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("unable to cast objectid")
	}
	hotel.ID = uid.Hex()

	return nil
}

func DeleteHotel(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := hotelsColl.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func UpdateHotel(ctx context.Context, hotel *types.Hotel) error {
	update := bson.M{
		"$set": bson.M{
			"name":     hotel.Name,
			"location": hotel.Location,
		},
	}

	oid, err := primitive.ObjectIDFromHex(hotel.ID)
	if err != nil {
		return err
	}

	_, err = hotelsColl.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return err
	}

	return nil
}
