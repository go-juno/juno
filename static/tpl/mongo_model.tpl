package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type {{.Class}} struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Add other fields here, e.g., Name string `bson:"name"`
	// This is a basic template, fields can be extended later
}
