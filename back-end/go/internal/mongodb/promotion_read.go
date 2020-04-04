package mongodb

import (
    "context"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) GetAllPromotions() ([]models.Promotion, error) {
    var cursor *mongo.Cursor
    var err error
    var promotion models.Promotion
    var promotions []models.Promotion

    cursor, err = d.DB.Collection("promotions").Find(
        context.Background(),
        bson.M{},
    )

    if err != nil {
        return promotions, err
    }

    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        promotion = models.Promotion{}

        err = cursor.Decode(&promotion)

        if err != nil {
            return promotions, err
        }

        promotions = append(promotions, promotion)
    }

    err = cursor.Err()

    if err != nil {
        return promotions, err
    }

    return promotions, nil
}

func (d *Datastore) GetPromotion(id string) (models.Promotion, error) {
    var objectID primitive.ObjectID
    var err error
    var promotion models.Promotion
    var result *mongo.SingleResult

    // It creates an ObjectID from a hex string.
    objectID, err = primitive.ObjectIDFromHex(id)

    if err != nil {
        return promotion, fmt.Errorf("the id isn't valid")
    }

    result = d.DB.Collection("promotions").FindOne(
        context.Background(),
        bson.M{"_id": objectID},
    )

    err = result.Err()

    if err != nil {
        if err != mongo.ErrNoDocuments {
            return promotion, err
        } else {
            return promotion, nil
        }
    }

    err = result.Decode(&promotion)

    if err != nil {
        return promotion, err
    }

    return promotion, nil
}
