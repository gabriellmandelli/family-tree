package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
