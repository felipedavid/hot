package storage

import (
	"context"
	"errors"

	"github.com/felipedavid/hot/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(ctx context.Context, id string) (*types.User, error) {
	var user types.User

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = userColl.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := userColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*types.User
	err = cur.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func InsertUser(ctx context.Context, user *types.User) error {
	res, err := userColl.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	uid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("unable to cast objectid")
	}
	user.ID = uid.Hex()

	return nil
}

var ErrNotFound = errors.New("not found")

func DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := userColl.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func UpdateUser(ctx context.Context, user *types.User) error {
	update := bson.D{
		{"$set", bson.D{
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"email", user.Email},
			{"hashed_password", user.HashedPassword},
		}},
	}

	oid, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	_, err = userColl.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return err
	}

	return nil
}
