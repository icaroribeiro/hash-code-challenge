package mongodb

import (
    "context"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) UpdateDiscountedDate(id string, discountedDate models.DiscountedDate) (int64, int64, error) {
    var objectID primitive.ObjectID
    var err error
    var singleResult *mongo.SingleResult
    var discountedDateAux models.DiscountedDate    
    var updateResult *mongo.UpdateResult

    // It creates an ObjectID from a hex string.
    objectID, err = primitive.ObjectIDFromHex(id)

    if err != nil {
        return 0, 0, fmt.Errorf("the id isn't valid")
    }

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
            return 0, 0, err
        }
    } else {
        err = singleResult.Decode(&discountedDateAux)

        if err != nil {
            return 0, 0, err
        }

        if objectID != discountedDateAux.ID {
            return 0, 0, fmt.Errorf("the discounted date with the id %s is already registered with the year %d, month %d and day %d", 
                discountedDateAux.ID.Hex(), discountedDateAux.Date.Year, discountedDateAux.Date.Month, discountedDateAux.Date.Day)
        }
    }

    updateResult, err = d.DB.Collection("discountedDates").UpdateOne(
        context.Background(),
        bson.M{"_id": objectID},
        bson.M{
            "$set": discountedDate,
        },
    )

    if err != nil {
        return 0, 0, err
    }

    return updateResult.MatchedCount, updateResult.ModifiedCount, nil
}
