package mongodb

import (
	"context"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) GetAllUsers() ([]models.User, error) {
	var cursor *mongo.Cursor
	var err error
	var user models.User
	var users []models.User

	cursor, err = d.DB.Collection("users").Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return users, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user = models.User{}

		err = cursor.Decode(&user)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	err = cursor.Err()

	if err != nil {
		return users, err
	}

	return users, nil
}

func (d *Datastore) GetUser(id string) (models.User, error) {
	var objectID primitive.ObjectID
	var err error
	var user models.User
	var result *mongo.SingleResult

	// It creates an ObjectID from a hex string.
	objectID, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return user, fmt.Errorf("the id isn't valid")
	}

	result = d.DB.Collection("users").FindOne(
		context.Background(),
		bson.M{"_id": objectID},
	)

	err = result.Err()

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return user, err
		} else {
			return user, nil
		}
	}

	err = result.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}
