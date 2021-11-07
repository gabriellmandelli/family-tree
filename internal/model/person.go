package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	Age           int                `json:"age" bson:"age"`
	RelationShips []RelationShip     `json:"relationships" bson:"relationships"`
}

type RelationShip struct {
	Person primitive.ObjectID `json:"person" bson:"person"`
	Type   string             `json:"type"`
}

type CreatePerson struct {
	Name          string         `json:"name" bson:"name"`
	Age           int            `json:"age" bson:"age"`
	RelationShips []RelationShip `json:"relationships" bson:"relationships"`
}

type UpdatePerson struct {
	Name          string         `json:"name" bson:"name"`
	Age           int            `json:"age" bson:"age"`
	RelationShips []RelationShip `json:"relationships" bson:"relationships"`
}
