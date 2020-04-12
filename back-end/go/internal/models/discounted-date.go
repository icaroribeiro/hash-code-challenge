package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type DiscountedDate struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    DiscountPct float32            `json:"discount_pct"`
    Date        Date               `json:"date"`
}
