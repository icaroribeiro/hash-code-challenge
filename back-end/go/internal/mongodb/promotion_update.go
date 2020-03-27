package mongodb

import (
	"context"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) UpdatePromotion(id string, promotion models.Promotion) (int64, int64, error) {
	var objectID primitive.ObjectID
	var err error
	var singleResult *mongo.SingleResult
	var promotionAux models.Promotion
	var updateResult *mongo.UpdateResult

	// Parsing a string id to MongoDB ObjectID.
	objectID, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, 0, fmt.Errorf("the id isn't valid")
	}

	singleResult = d.DB.Collection("promotions").FindOne(
		context.Background(),
		bson.M{"code": promotion.Code},
	)

	err = singleResult.Err()

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return 0, 0, err
		}
	} else {
		err = singleResult.Decode(&promotionAux)

		if err != nil {
			return 0, 0, err
		}
	
		if objectID != promotionAux.ID {
			return 0, 0, fmt.Errorf("the promotion with the id %s is already registered with the code %s", 
				promotionAux.ID.Hex(), promotion.Code)
		}
	}

	if len(promotion.Products) == 0 {
		updateResult, err = d.DB.Collection("promotions").UpdateOne(
			context.Background(),
			bson.M{"_id": objectID},
			bson.M{
				"$set": bson.M{
					"code":             promotion.Code,
					"title":            promotion.Title,
					"description":      promotion.Description,
					"max_discount_pct": promotion.MaxDiscountPct,
				},
				"$unset": bson.M{
					"products": "",
				},
			},
		)
	} else {
		updateResult, err = d.DB.Collection("promotions").UpdateOne(
			context.Background(),
			bson.M{"_id": objectID},
			bson.M{
				"$set": bson.M{
					"code":             promotion.Code,
					"title":            promotion.Title,
					"description":      promotion.Description,
					"max_discount_pct": promotion.MaxDiscountPct,
					"products":         promotion.Products,
				},
			},
		)
	}

	if err != nil {
		return 0, 0, err
	}

	return updateResult.MatchedCount, updateResult.ModifiedCount, nil
}