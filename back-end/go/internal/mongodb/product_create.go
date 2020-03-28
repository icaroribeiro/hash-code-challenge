package mongodb

import (
	"context"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Datastore) CreateProduct(product models.Product) (models.Product, error) {
	var result *mongo.InsertOneResult
	var err error
	var id primitive.ObjectID
	var isOK bool

	result, err = d.DB.Collection("products").InsertOne(
		context.Background(),
		product,
	)

	if err != nil {
		return product, err
	}

	// Type assertion.
	id, isOK = result.InsertedID.(primitive.ObjectID)

	if isOK {
		product.ID = id
	} else {
		return product, fmt.Errorf("it wasn't possible to obtain the id of generated document")
	}

	return product, nil
}
