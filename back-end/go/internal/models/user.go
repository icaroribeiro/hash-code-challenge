package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    FirstName   string             `json:"first_name"`
    LastName    string             `json:"last_name"`
    DateOfBirth Date               `json:"date_of_birth"`
}
