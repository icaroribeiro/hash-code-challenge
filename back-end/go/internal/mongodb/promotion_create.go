package mongodb

import (
	"context"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) CreatePromotion(promotion models.Promotion) (models.Promotion, error) {
	var singleResult *mongo.SingleResult
	var err error
	var promotionAux models.Promotion
	var insertOneResult *mongo.InsertOneResult
	var id primitive.ObjectID
	var isOK bool

	singleResult = d.DB.Collection("promotions").FindOne(
		context.Background(),
		bson.M{"code": promotion.Code},
	)

	err = singleResult.Err()

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return promotion, err
		}
	} else {
		err = singleResult.Decode(&promotionAux)

		if err != nil {
			return promotion, err
		}

		return promotion, fmt.Errorf("the promotion with the id %s is already registered with the code %s",
			promotionAux.ID.Hex(), promotionAux.Code)
	}

	insertOneResult, err = d.DB.Collection("promotions").InsertOne(
		context.Background(),
		promotion,
	)

	if err != nil {
		return promotion, err
	}

	// Type assertion.
	id, isOK = insertOneResult.InsertedID.(primitive.ObjectID)

	if isOK {
		promotion.ID = id
	} else {
		return promotion, fmt.Errorf("it wasn't possible to get the id of the generated document")
	}

	return promotion, nil
}
