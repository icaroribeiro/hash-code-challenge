package mongodb

import (
	"context"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) GetAllDiscountedDates() ([]models.DiscountedDate, error) {
	var cursor *mongo.Cursor
	var err error
	var discountedDate models.DiscountedDate
	var discountedDates []models.DiscountedDate

	cursor, err = d.DB.Collection("discountedDates").Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return discountedDates, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		discountedDate = models.DiscountedDate{}

		err = cursor.Decode(&discountedDate)

		if err != nil {
			return discountedDates, err
		}

		discountedDates = append(discountedDates, discountedDate)
	}

	err = cursor.Err()

	if err != nil {
		return discountedDates, err
	}

	return discountedDates, nil
}

func (d *Datastore) GetDiscountedDate(id string) (models.DiscountedDate, error) {
	var objectID primitive.ObjectID
	var err error
	var discountedDate models.DiscountedDate
	var result *mongo.SingleResult

	// It creates an ObjectID from a hex string.
	objectID, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return discountedDate, fmt.Errorf("the id isn't valid")
	}

	result = d.DB.Collection("discountedDates").FindOne(
		context.Background(),
		bson.M{"_id": objectID},
	)

	err = result.Err()

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return discountedDate, err
		} else {
			return discountedDate, nil
		}
	}

	err = result.Decode(&discountedDate)

	if err != nil {
		return discountedDate, err
	}

	return discountedDate, nil
}
