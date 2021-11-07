package database

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/model"
	"github.com/gabriellmandelli/family-tree/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cntx context.Context

type PersonRepository interface {
	Count() (int64, error)
	FindAll(name string, limit int64, offset int64) ([]model.Person, error)
	FindById(id string) (*model.Person, error)
	Save(payload *model.Person) (*model.Person, error)
	Update(id string, payload *model.Person) (*model.Person, error)
	Delete(id string) error
}

type PersonRepositoryImpl struct {
	Connection *mongo.Database
}

func NewPersonRepository(Connection *mongo.Database) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{Connection: Connection}
}

func (pr *PersonRepositoryImpl) Count() (int64, error) {
	countRecord, errorHandlerCount := pr.Connection.Collection("person").CountDocuments(cntx, bson.M{}, nil)

	if !util.GlobalErrorDatabaseException(errorHandlerCount) {
		return 0, errorHandlerCount
	}

	return countRecord, nil
}

func (pr *PersonRepositoryImpl) FindAll(name string, limit int64, offset int64) ([]model.Person, error) {
	var (
		person        model.Person
		Person        []model.Person
		filterOptions = options.Find()
		csr           *mongo.Cursor
		errorCsr      error
	)

	filterOptions.SetLimit(limit)
	filterOptions.SetSkip(offset)

	if name != "" {
		csr, errorCsr = pr.Connection.Collection("person").Find(cntx, bson.M{"name": name}, filterOptions)
		if !util.GlobalErrorDatabaseException(errorCsr) {
			return nil, errorCsr
		}
	} else {
		csr, errorCsr = pr.Connection.Collection("person").Find(cntx, bson.M{}, filterOptions)

		if !util.GlobalErrorDatabaseException(errorCsr) {
			return nil, errorCsr
		}
	}

	for csr.Next(cntx) {
		errorHandlerDecodeData := csr.Decode(&person)

		if !util.GlobalErrorDatabaseException(errorHandlerDecodeData) {
			return nil, errorHandlerDecodeData
		}

		Person = append(Person, person)
	}

	return Person, nil
}

func (pr *PersonRepositoryImpl) FindById(id string) (*model.Person, error) {
	var (
		person      model.Person
		personId, _ = primitive.ObjectIDFromHex(id)
		filter      = bson.M{"_id": personId}
	)

	errorGetOneperson := pr.Connection.Collection("person").FindOne(cntx, filter).Decode(&person)

	if !util.GlobalErrorDatabaseException(errorGetOneperson) {
		return nil, errorGetOneperson
	}

	return &person, nil

}

func (pr *PersonRepositoryImpl) Save(payload *model.Person) (*model.Person, error) {
	payload.ID = primitive.NewObjectID()
	_, errorHandlerSaveperson := pr.Connection.Collection("person").InsertOne(cntx, payload)

	if !util.GlobalErrorDatabaseException(errorHandlerSaveperson) {
		return nil, errorHandlerSaveperson
	}

	return payload, nil
}

func (pr *PersonRepositoryImpl) Update(id string, payload *model.Person) (*model.Person, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objectID,
	}

	updateField := bson.M{
		"$set": bson.M{
			"name":          payload.Name,
			"age":           payload.Age,
			"relationships": payload.RelationShips,
		}}

	_, errorHandlerUpdatePerson := pr.Connection.Collection("person").UpdateOne(cntx, filter, updateField)

	if !util.GlobalErrorDatabaseException(errorHandlerUpdatePerson) {
		return nil, errorHandlerUpdatePerson
	}

	return payload, nil
}

func (pr *PersonRepositoryImpl) Delete(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}

	_, errorHandlerDelete := pr.Connection.Collection("person").DeleteOne(cntx, filter)

	if !util.GlobalErrorDatabaseException(errorHandlerDelete) {
		return errorHandlerDelete
	}

	return nil
}
