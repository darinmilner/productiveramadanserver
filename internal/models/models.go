package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Signup struct {
	FirstName string
	LastName  string
	Email     string
}

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
}

type HijriDate struct {
	Day   int
	Month string
	Year  int
}
