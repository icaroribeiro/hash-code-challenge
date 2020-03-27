package models

type Date struct {
	Year  int `json:"year" bson:"year"`
	Month int `json:"month" bson:"month"`
	Day   int `json:"day" bson:"day"`
}
