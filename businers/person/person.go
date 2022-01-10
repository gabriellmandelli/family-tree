package person

import (
	"context"

	"github.com/gabriellmandelli/family-tree/adapter/config"
	"github.com/gabriellmandelli/family-tree/adapter/database"
	"github.com/joomcode/errorx"
)

type PersonService interface {
	FindInBatch(ctx context.Context, personIds []string) ([]Person, *errorx.Error)
	AddPerson(ctx context.Context, person *Person) (*Person, *errorx.Error)
	FindAllPerson(ctx context.Context, name string) ([]Person, *errorx.Error)
}

type PersonServiceImpl struct {
	personRepository *PersonRepositoryImpl
}

func NewPersonService() PersonService {

	db, _ := database.NewMongoDbClient(context.TODO(), config.GetConfig())
	personRespository := NewPersonRepository(db)

	return &PersonServiceImpl{
		personRepository: personRespository,
	}
}
func (p *PersonServiceImpl) FindInBatch(ctx context.Context, personIds []string) ([]Person, *errorx.Error) {
	return p.personRepository.FindInBatch(ctx, personIds)
}

func (p *PersonServiceImpl) FindAllPerson(ctx context.Context, name string) ([]Person, *errorx.Error) {
	return p.personRepository.FindAll(name, 0, 0)
}

func (p *PersonServiceImpl) AddPerson(ctx context.Context, person *Person) (*Person, *errorx.Error) {
	return p.personRepository.Save(person)
}
