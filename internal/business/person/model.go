package person

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID   string `json:"id,omitempty" `
	Name string             `json:"name"`
}
