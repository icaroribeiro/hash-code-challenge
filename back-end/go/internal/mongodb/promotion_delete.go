package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) DeletePromotion(id string) (int64, error) {
	var objectID primitive.ObjectID
	var err error
	var result *mongo.DeleteResult

	// Parsing a string id to MongoDB ObjectID.
	objectID, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, fmt.Errorf("the id isn't valid")
	}

	result, err = d.DB.Collection("promotions").DeleteOne(
		context.Background(),
		bson.M{"_id": objectID},
	)

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
