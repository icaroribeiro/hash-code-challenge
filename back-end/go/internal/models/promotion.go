package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Promotion struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Code           string             `json:"code" bson:"code"`
	Title          string             `json:"title" bson:"title"`
	Description    string             `json:"description" bson:"description"`
	MaxDiscountPct float32            `json:"max_discount_pct" bson:"max_discount_pct"`
	Products       []string           `json:"products,omitempty" bson:"products,omitempty"`
}
