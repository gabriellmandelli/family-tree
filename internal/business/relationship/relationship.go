package relationship

import (
	"context"

	"github.com/joomcode/errorx"
)

type RelationShipService interface {
	Add(ctx context.Context, person *RelationShip) (*RelationShip, *errorx.Error)
	FindAll(ctx context.Context, parentID string, childrenID string) ([]RelationShip, *errorx.Error)
}

type RelationShipServiceImpl struct {
	relationShipRepository RelationShipRepository
}

func NewRelationShipService(r RelationShipRepository) RelationShipService {
	return &RelationShipServiceImpl{
		relationShipRepository: r,
	}
}

func (p *RelationShipServiceImpl) FindAll(ctx context.Context, parentID string, childrenID string) ([]RelationShip, *errorx.Error) {
	return p.relationShipRepository.FindAll(parentID, childrenID, 0, 0)
}

func (p *RelationShipServiceImpl) Add(ctx context.Context, person *RelationShip) (*RelationShip, *errorx.Error) {
	return p.relationShipRepository.Save(person)
}
