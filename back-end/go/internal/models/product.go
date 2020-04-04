package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
    ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    PriceInCents int                `json:"price_in_cents" bson:"price_in_cents"`
    Title        string             `json:"title" bson:"title"`
    Description  string             `json:"description" bson:"description"`
}
