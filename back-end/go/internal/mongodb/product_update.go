package mongodb

import (
    "context"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) UpdateProduct(id string, product models.Product) (int64, int64, error) {
    var objectID primitive.ObjectID
    var err error
    var result *mongo.UpdateResult

    // It creates an ObjectID from a hex string.
    objectID, err = primitive.ObjectIDFromHex(id)

    if err != nil {
        return 0, 0, fmt.Errorf("the id isn't valid")
    }

    result, err = d.DB.Collection("products").UpdateOne(
        context.Background(),
        bson.M{"_id": objectID},
        bson.M{
            "$set": product,
        },
    )

    if err != nil {
        return 0, 0, err
    }

    return result.MatchedCount, result.ModifiedCount, nil
}
