package service

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/model"
	"github.com/joomcode/errorx"
)

type FamilyTreeService interface {
	GetFamilyTree(ctx *context.Context, personID string) (*model.FamilyTree, *errorx.Error)
}

type FamilyTreeServiceImpl struct {
	personService       PersonService
	relationShipService RelationShipService
}

func NewFamilyTreeService() FamilyTreeService {
	return &FamilyTreeServiceImpl{
		personService:       NewPersonService(),
		relationShipService: NewRelationShipService(),
	}
}

func (ft *FamilyTreeServiceImpl) GetFamilyTree(ctx *context.Context, personID string) (*model.FamilyTree, *errorx.Error) {
	familyTree := model.FamilyTree{}
	queryPerson := []string{}

	relationsShips, errx := ft.relationShipService.FindAll(ctx, "", "")

	if errx != nil {
		return &familyTree, errx
	}

	ft.parentsOfParents(&familyTree.Members, personID, queryPerson, relationsShips)

	personData, errx := ft.personService.FindInBatch(ctx, queryPerson)

	if errx != nil {
		return &familyTree, errx
	}

	ft.updatePersonInfo(&familyTree.Members, personData)

	return &familyTree, nil
}

func (ft *FamilyTreeServiceImpl) parentsOfParents(member *model.Members, childrenID string, queryPerson []string, relationsShips []model.RelationShip) {
	member.ID = childrenID
	member.Parents = make([]model.Members, 0)
	member.Childrens = make([]model.Members, 0)
	queryPerson = append(queryPerson, childrenID)
	for _, relation := range relationsShips {
		if childrenID == relation.ChildrenID.Hex() {
			parent := model.Members{}
			parent.ID = relation.ParentID.Hex()
			ft.parentsOfParents(&parent, parent.ID, queryPerson, relationsShips)
			member.Parents = append(member.Parents, parent)
		}
	}
}

func (ft *FamilyTreeServiceImpl) updatePersonInfo(members *model.Members, personData []model.Person) {
	for i := range personData {
		if members.ID == personData[i].ID.Hex() {
			members.Name = personData[i].Name
		}
	}
	for i := range members.Parents {
		ft.updatePersonInfo(&members.Parents[i], personData)
	}
}
