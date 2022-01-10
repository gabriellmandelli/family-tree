package relationship

import (
	"context"

	"github.com/gabriellmandelli/family-tree/adapter/config"
	"github.com/gabriellmandelli/family-tree/adapter/database"
	"github.com/joomcode/errorx"
)

type RelationShipService interface {
	Add(ctx context.Context, person *RelationShip) (*RelationShip, *errorx.Error)
	FindAll(ctx context.Context, parentID string, childrenID string) ([]RelationShip, *errorx.Error)
}

type RelationShipServiceImpl struct {
	relationShipRepository *RelationShipRepositoryImpl
}

func NewRelationShipService() RelationShipService {
	db, _ := database.NewMongoDbClient(context.TODO(), config.GetConfig())
	rsRespository := NewRelationShipRepository(db)
	return &RelationShipServiceImpl{
		relationShipRepository: rsRespository,
	}
}

func (p *RelationShipServiceImpl) FindAll(ctx context.Context, parentID string, childrenID string) ([]RelationShip, *errorx.Error) {
	return p.relationShipRepository.FindAll(parentID, childrenID, 0, 0)
}

func (p *RelationShipServiceImpl) Add(ctx context.Context, person *RelationShip) (*RelationShip, *errorx.Error) {
	return p.relationShipRepository.Save(person)
}
