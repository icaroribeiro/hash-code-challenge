package mongodb

import (
    "context"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) GetAllProducts() ([]models.Product, error) {
    var cursor *mongo.Cursor
    var err error
    var product models.Product
    var products []models.Product

    cursor, err = d.DB.Collection("products").Find(
        context.Background(),
        bson.M{},
    )

    if err != nil {
        return products, err
    }

    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        product = models.Product{}

        err = cursor.Decode(&product)

        if err != nil {
            return products, err
        }

        products = append(products, product)
    }

    err = cursor.Err()

    if err != nil {
        return products, err
    }

    return products, nil
}

func (d *Datastore) GetProduct(id string) (models.Product, error) {
    var objectID primitive.ObjectID
    var err error
    var product models.Product
    var result *mongo.SingleResult

    // It creates an ObjectID from a hex string.
    objectID, err = primitive.ObjectIDFromHex(id)

    if err != nil {
        return product, fmt.Errorf("the id isn't valid")
    }

    result = d.DB.Collection("products").FindOne(
        context.Background(),
        bson.M{"_id": objectID},
    )

    err = result.Err()

    if err != nil {
        if err != mongo.ErrNoDocuments {
            return product, err
        } else {
            return product, nil
        }
    }

    err = result.Decode(&product)

    if err != nil {
        return product, err
    }

    return product, nil
}
