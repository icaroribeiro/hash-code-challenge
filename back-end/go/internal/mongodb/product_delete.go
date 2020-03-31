package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) DeleteProduct(id string) (int64, error) {
	var objectID primitive.ObjectID
	var err error
	var result *mongo.DeleteResult

	// It creates an ObjectID from a hex string.
	objectID, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, fmt.Errorf("the id isn't valid")
	}

	result, err = d.DB.Collection("products").DeleteOne(
		context.Background(),
		bson.M{"_id": objectID},
	)

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
