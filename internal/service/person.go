package service

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/config"
	"github.com/gabriellmandelli/family-tree/internal/database"
	"github.com/gabriellmandelli/family-tree/internal/model"
	"github.com/joomcode/errorx"
)

type PersonService interface {
	AddPerson(ctx *context.Context, person *model.Person) (*model.Person, *errorx.Error)
	FindAllPerson(ctx *context.Context, name string) ([]model.Person, *errorx.Error)
}

type PersonServiceImpl struct {
	personRepository *database.PersonRepositoryImpl
}

func NewPersonService() PersonService {

	db, _ := database.NewMongoDbClient(context.TODO(), config.GetConfig())
	personRespository := database.NewPersonRepository(db)

	return &PersonServiceImpl{
		personRepository: personRespository,
	}
}

func (p *PersonServiceImpl) FindAllPerson(ctx *context.Context, name string) ([]model.Person, *errorx.Error) {
	persons, err := p.personRepository.FindAll(name, 0, 0)

	if err != nil {
		return nil, errorx.Decorate(err, "Error to find all")
	}

	return persons, nil
}

func (p *PersonServiceImpl) AddPerson(ctx *context.Context, person *model.Person) (*model.Person, *errorx.Error) {

	response, err := p.personRepository.Save(person)

	if err != nil {
		return nil, errorx.Decorate(err, "Error to add person")
	}

	return response, nil
}
