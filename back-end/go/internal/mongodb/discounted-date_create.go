package mongodb

import (
	"context"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) CreateDiscountedDate(discountedDate models.DiscountedDate) (models.DiscountedDate, error) {
	var singleResult *mongo.SingleResult
	var err error
	var discountedDateAux models.DiscountedDate
	var insertOneResult *mongo.InsertOneResult
	var id primitive.ObjectID
	var isOK bool

	singleResult = d.DB.Collection("discountedDates").FindOne(
		context.Background(),
		bson.M{
			"date.year": discountedDate.Date.Year,
			"date.month": discountedDate.Date.Month,
			"date.day": discountedDate.Date.Day,
		},
	)

	err = singleResult.Err()

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return discountedDate, err
		}
	} else {
		err = singleResult.Decode(&discountedDateAux)

		if err != nil {
			return discountedDate, err
		}

		return discountedDate, fmt.Errorf("the discounted date with the id %s is already registered with the year %d, month %d and day %d", 
			discountedDateAux.ID.Hex(), discountedDateAux.Date.Year, discountedDateAux.Date.Month, discountedDateAux.Date.Day)
	}

	insertOneResult, err = d.DB.Collection("discountedDates").InsertOne(
		context.Background(),
		discountedDate,
	)

	if err != nil {
		return discountedDate, err
	}

	// Type assertion.
	id, isOK = insertOneResult.InsertedID.(primitive.ObjectID)

	if isOK {
		discountedDate.ID = id
	} else {
		return discountedDate, fmt.Errorf("it wasn't possible to get the id of generated document")
	}

	return discountedDate, nil
}
