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
	ft.childrenOfChildres(&familyTree.Members, personID, queryPerson, relationsShips)

	personData, errx := ft.personService.FindInBatch(ctx, queryPerson)

	if errx != nil {
		return &familyTree, errx
	}

	ft.updatePersonInfo(&familyTree.Members, personData)

	return &familyTree, nil
}

func (ft *FamilyTreeServiceImpl) parentsOfParents(member *model.Members, memberID string, queryPerson []string, relationsShips []model.RelationShip) {
	member.ID = memberID
	member.Parents = make([]model.Members, 0)
	queryPerson = append(queryPerson, memberID)
	for _, relation := range relationsShips {
		if memberID == relation.ChildrenID.Hex() {
			parent := model.Members{}
			parent.ID = relation.ParentID.Hex()
			ft.parentsOfParents(&parent, parent.ID, queryPerson, relationsShips)
			member.Parents = append(member.Parents, parent)
		}
	}
}

func (ft *FamilyTreeServiceImpl) childrenOfChildres(member *model.Members, memberID string, queryPerson []string, relationsShips []model.RelationShip) {
	member.ID = memberID
	member.Childrens = make([]model.Members, 0)
	queryPerson = append(queryPerson, memberID)
	for _, relation := range relationsShips {
		if memberID == relation.ParentID.Hex() {
			children := model.Members{}
			children.ID = relation.ChildrenID.Hex()
			ft.childrenOfChildres(&children, children.ID, queryPerson, relationsShips)
			member.Childrens = append(member.Childrens, children)
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

	for i := range members.Childrens {
		ft.updatePersonInfo(&members.Childrens[i], personData)
	}
}
