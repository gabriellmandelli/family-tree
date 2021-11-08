package service

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/config"
	"github.com/gabriellmandelli/family-tree/internal/database"
	"github.com/gabriellmandelli/family-tree/internal/model"
	"github.com/joomcode/errorx"
)

type RelationShipService interface {
	Add(ctx *context.Context, person *model.RelationShip) (*model.RelationShip, *errorx.Error)
	FindAll(ctx *context.Context, parentID string, childrenID string) ([]model.RelationShip, *errorx.Error)
}

type RelationShipServiceImpl struct {
	relationShipRepository *database.RelationShipRepositoryImpl
}

func NewRelationShipService() RelationShipService {

	db, _ := database.NewMongoDbClient(context.TODO(), config.GetConfig())
	rsRespository := database.NewRelationShipRepository(db)

	return &RelationShipServiceImpl{
		relationShipRepository: rsRespository,
	}
}

func (p *RelationShipServiceImpl) FindAll(ctx *context.Context, parentID string, childrenID string) ([]model.RelationShip, *errorx.Error) {
	return p.relationShipRepository.FindAll(parentID, childrenID, 0, 0)
}

func (p *RelationShipServiceImpl) Add(ctx *context.Context, person *model.RelationShip) (*model.RelationShip, *errorx.Error) {
	return p.relationShipRepository.Save(person)
}
