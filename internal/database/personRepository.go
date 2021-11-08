package database

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/model"
	"github.com/joomcode/errorx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	personCollection = "person"
)

var cntx context.Context

type PersonRepository interface {
	Count() (int64, *errorx.Error)
	FindInBatch(ctx *context.Context, personIDs []string) ([]model.Person, *errorx.Error)
	FindAll(name string, limit int64, offset int64) ([]model.Person, *errorx.Error)
	FindById(id string) (*model.Person, *errorx.Error)
	Save(payload *model.Person) (*model.Person, *errorx.Error)
	Update(id string, payload *model.Person) (*model.Person, *errorx.Error)
	Delete(id string) *errorx.Error
}

type PersonRepositoryImpl struct {
	Connection *mongo.Database
}

func NewPersonRepository(Connection *mongo.Database) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{Connection: Connection}
}

func (pr *PersonRepositoryImpl) Count() (int64, *errorx.Error) {
	countRecord, errorHandlerCount := pr.Connection.Collection(personCollection).CountDocuments(cntx, bson.M{}, nil)

	if errorHandlerCount != nil {
		return 0, errorx.Decorate(errorHandlerCount, "Database error")
	}

	return countRecord, nil
}

func (pr *PersonRepositoryImpl) FindInBatch(ctx *context.Context, personIDs []string) ([]model.Person, *errorx.Error) {
	var (
		person        model.Person
		persons       []model.Person
		filterOptions = options.Find()
		csr           *mongo.Cursor
		errorCsr      error
	)

	if len(personIDs) > 0 {

		var filter []bson.M

		for i := range personIDs {
			filter = append(filter, bson.M{"_id": personIDs[i]})
		}

		csr, errorCsr = pr.Connection.Collection(personCollection).Find(*ctx, filter, filterOptions)
		if errorCsr != nil {
			return nil, errorx.Decorate(errorCsr, "Database error")
		}
	} else {
		csr, errorCsr = pr.Connection.Collection(personCollection).Find(*ctx, bson.M{}, filterOptions)
		if errorCsr != nil {
			return nil, errorx.Decorate(errorCsr, "Database error")
		}
	}

	for csr.Next(*ctx) {
		errorHandlerDecodeData := csr.Decode(&person)

		if errorHandlerDecodeData != nil {
			return nil, errorx.Decorate(errorHandlerDecodeData, "Database error")
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (pr *PersonRepositoryImpl) FindAll(name string, limit int64, offset int64) ([]model.Person, *errorx.Error) {
	var (
		person        model.Person
		persons       []model.Person
		filterOptions = options.Find()
		csr           *mongo.Cursor
		errorCsr      error
	)

	filterOptions.SetLimit(limit)
	filterOptions.SetSkip(offset)

	if name != "" {
		csr, errorCsr = pr.Connection.Collection(personCollection).Find(cntx, bson.M{"name": name}, filterOptions)
		if errorCsr != nil {
			return nil, errorx.Decorate(errorCsr, "Database error")
		}
	} else {
		csr, errorCsr = pr.Connection.Collection(personCollection).Find(cntx, bson.M{}, filterOptions)
		if errorCsr != nil {
			return nil, errorx.Decorate(errorCsr, "Database error")
		}
	}

	for csr.Next(cntx) {
		errorHandlerDecodeData := csr.Decode(&person)

		if errorHandlerDecodeData != nil {
			return nil, errorx.Decorate(errorHandlerDecodeData, "Database error")
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (pr *PersonRepositoryImpl) FindById(id string) (*model.Person, *errorx.Error) {
	var (
		person      model.Person
		personId, _ = primitive.ObjectIDFromHex(id)
		filter      = bson.M{"_id": personId}
	)

	errorGetOneperson := pr.Connection.Collection(personCollection).FindOne(cntx, filter).Decode(&person)

	if errorGetOneperson != nil {
		return nil, errorx.Decorate(errorGetOneperson, "Database error")
	}

	return &person, nil

}

func (pr *PersonRepositoryImpl) Save(payload *model.Person) (*model.Person, *errorx.Error) {
	payload.ID = primitive.NewObjectID()
	_, errorHandlerSaveperson := pr.Connection.Collection(personCollection).InsertOne(cntx, payload)

	if errorHandlerSaveperson != nil {
		return nil, errorx.Decorate(errorHandlerSaveperson, "Database error")
	}

	return payload, nil
}

func (pr *PersonRepositoryImpl) Update(id string, payload *model.Person) (*model.Person, *errorx.Error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objectID,
	}

	updateField := bson.M{
		"$set": bson.M{
			"name": payload.Name,
		}}

	_, errorHandlerUpdatePerson := pr.Connection.Collection(personCollection).UpdateOne(cntx, filter, updateField)

	if errorHandlerUpdatePerson != nil {
		return nil, errorx.Decorate(errorHandlerUpdatePerson, "Database error")
	}

	return payload, nil
}

func (pr *PersonRepositoryImpl) Delete(id string) *errorx.Error {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}

	_, errorHandlerDelete := pr.Connection.Collection(personCollection).DeleteOne(cntx, filter)

	if errorHandlerDelete != nil {
		return errorx.Decorate(errorHandlerDelete, "Database error")
	}

	return nil
}
