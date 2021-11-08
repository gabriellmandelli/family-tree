package service

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/config"
	"github.com/gabriellmandelli/family-tree/internal/database"
	"github.com/gabriellmandelli/family-tree/internal/model"
	"github.com/joomcode/errorx"
)

type PersonService interface {
	FindInBatch(ctx *context.Context, personIds []string) ([]model.Person, *errorx.Error)
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
func (p *PersonServiceImpl) FindInBatch(ctx *context.Context, personIds []string) ([]model.Person, *errorx.Error) {
	return p.personRepository.FindInBatch(ctx, personIds)
}

func (p *PersonServiceImpl) FindAllPerson(ctx *context.Context, name string) ([]model.Person, *errorx.Error) {
	return p.personRepository.FindAll(name, 0, 0)
}

func (p *PersonServiceImpl) AddPerson(ctx *context.Context, person *model.Person) (*model.Person, *errorx.Error) {
	return p.personRepository.Save(person)
}
