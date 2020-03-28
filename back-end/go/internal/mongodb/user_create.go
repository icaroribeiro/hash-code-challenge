package mongodb

import (
	"context"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) CreateUser(user models.User) (models.User, error) {
	var result *mongo.InsertOneResult
	var err error
	var id primitive.ObjectID
	var isOK bool

	result, err = d.DB.Collection("users").InsertOne(
		context.Background(),
		user,
	)

	if err != nil {
		return user, err
	}

	// Type assertion.
	id, isOK = result.InsertedID.(primitive.ObjectID)

	if isOK {
		user.ID = id
	} else {
		return user, fmt.Errorf("it wasn't possible to get the id of the generated document")
	}

	return user, nil
}
