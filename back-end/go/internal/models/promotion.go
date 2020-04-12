package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Promotion struct {
    ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Code           string             `json:"code"`
    Title          string             `json:"title"`
    Description    string             `json:"description"`
    MaxDiscountPct float32            `json:"max_discount_pct"`
    Products       []string           `json:"products,omitempty"`
}
