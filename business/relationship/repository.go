package relationship

import (
	"context"

	"github.com/joomcode/errorx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	relationShipCollection = "relationship"
)

type RelationShipRepository interface {
	Count() (int64, *errorx.Error)
	FindAll(parentID string, childrenID string, limit int64, offset int64) ([]RelationShip, *errorx.Error)
	FindById(id string) (*RelationShip, *errorx.Error)
	Save(payload *RelationShip) (*RelationShip, *errorx.Error)
	Update(id string, payload *RelationShip) (*RelationShip, *errorx.Error)
	Delete(id string) *errorx.Error
}

type RelationShipRepositoryImpl struct {
	Connection *mongo.Database
}

func NewRelationShipRepository(Connection *mongo.Database) RelationShipRepository {
	return &RelationShipRepositoryImpl{Connection: Connection}
}

func (pr *RelationShipRepositoryImpl) Count() (int64, *errorx.Error) {
	countRecord, err := pr.Connection.Collection(relationShipCollection).CountDocuments(context.TODO(), bson.M{}, nil)

	if err != nil {
		return 0, errorx.Decorate(err, "Database error")
	}

	return countRecord, nil
}

func (pr *RelationShipRepositoryImpl) FindAll(parentID string, childrenID string, limit int64, offset int64) ([]RelationShip, *errorx.Error) {
	var (
		person        RelationShip
		persons       []RelationShip
		filterOptions = options.Find()
		csr           *mongo.Cursor
		err           error
		query         bson.M
	)

	filterOptions.SetLimit(limit)
	filterOptions.SetSkip(offset)

	switch true {
	case parentID != "" && childrenID != "":
		query = bson.M{"_id": parentID, "children_id": childrenID}
	case parentID != "":
		query = bson.M{"_id": parentID}
	case childrenID != "":
		query = bson.M{"children_id": childrenID}
	default:
		query = bson.M{}
	}

	csr, err = pr.Connection.Collection(relationShipCollection).Find(context.TODO(), query, filterOptions)

	if err != nil {
		return nil, errorx.Decorate(err, "Database error")
	}

	for csr.Next(context.TODO()) {
		err := csr.Decode(&person)

		if err != nil {
			return nil, errorx.Decorate(err, "Database error")
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (pr *RelationShipRepositoryImpl) FindById(id string) (*RelationShip, *errorx.Error) {
	var (
		person      RelationShip
		personId, _ = primitive.ObjectIDFromHex(id)
		filter      = bson.M{"_id": personId}
	)

	err := pr.Connection.Collection(relationShipCollection).FindOne(context.TODO(), filter).Decode(&person)

	if err != nil {
		return nil, errorx.Decorate(err, "Database error")
	}

	return &person, nil
}

func (pr *RelationShipRepositoryImpl) Save(payload *RelationShip) (*RelationShip, *errorx.Error) {
	_, err := pr.Connection.Collection(relationShipCollection).InsertOne(context.TODO(), payload)

	if err != nil {
		return nil, errorx.Decorate(err, "Database error")
	}

	return payload, nil
}

func (pr *RelationShipRepositoryImpl) Update(id string, payload *RelationShip) (*RelationShip, *errorx.Error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objectID,
	}

	updateField := bson.M{
		"$set": bson.M{
			"_id":         payload.ParentID,
			"children_id": payload.ChildrenID,
		}}

	_, err := pr.Connection.Collection(relationShipCollection).UpdateOne(context.TODO(), filter, updateField)

	if err != nil {
		return nil, errorx.Decorate(err, "Database error")
	}

	return payload, nil
}

func (pr *RelationShipRepositoryImpl) Delete(id string) *errorx.Error {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}

	_, err := pr.Connection.Collection(relationShipCollection).DeleteOne(context.TODO(), filter)

	if err != nil {
		return errorx.Decorate(err, "Database error")
	}

	return nil
}
