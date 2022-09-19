package relationship

import "go.mongodb.org/mongo-driver/bson/primitive"

type RelationShip struct {
	RelationID primitive.ObjectID `json:"id" bson:"_id"`
	ParentID   primitive.ObjectID `json:"parent_id" bson:"parent_id"`
	ChildrenID primitive.ObjectID `json:"children_id" bson:"children_id"`
}
