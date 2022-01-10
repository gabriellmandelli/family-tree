package relationship

import "go.mongodb.org/mongo-driver/bson/primitive"

type RelationShip struct {
	ParentID   primitive.ObjectID `json:"parent_id" bson:"_id"`
	ChildrenID primitive.ObjectID `json:"children_id" bson:"children_id"`
}
