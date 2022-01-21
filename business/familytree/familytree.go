package familytree

import (
	"context"

	"github.com/gabriellmandelli/family-tree/business/person"
	"github.com/gabriellmandelli/family-tree/business/relationship"
	"github.com/joomcode/errorx"
)

type FamilyTreeService interface {
	GetFamilyTree(ctx context.Context, personID string) (*FamilyTree, *errorx.Error)
}

type FamilyTreeServiceImpl struct {
	personService       person.PersonService
	relationShipService relationship.RelationShipService
}

func NewFamilyTreeService(personService person.PersonService, relationShipService relationship.RelationShipService) FamilyTreeService {
	return &FamilyTreeServiceImpl{
		personService:       personService,
		relationShipService: relationShipService,
	}
}

func (ft *FamilyTreeServiceImpl) GetFamilyTree(ctx context.Context, personID string) (*FamilyTree, *errorx.Error) {
	familyTree := FamilyTree{}
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

func (ft *FamilyTreeServiceImpl) parentsOfParents(member *Members, memberID string, queryPerson []string, relationsShips []relationship.RelationShip) {
	member.ID = memberID
	member.Parents = make([]Members, 0)
	queryPerson = append(queryPerson, memberID)
	for _, relation := range relationsShips {
		if memberID == relation.ChildrenID.Hex() {
			parent := Members{}
			parent.ID = relation.ParentID.Hex()
			ft.parentsOfParents(&parent, parent.ID, queryPerson, relationsShips)
			member.Parents = append(member.Parents, parent)
		}
	}
}

func (ft *FamilyTreeServiceImpl) childrenOfChildres(member *Members, memberID string, queryPerson []string, relationsShips []relationship.RelationShip) {
	member.ID = memberID
	member.Childrens = make([]Members, 0)
	queryPerson = append(queryPerson, memberID)
	for _, relation := range relationsShips {
		if memberID == relation.ParentID.Hex() {
			children := Members{}
			children.ID = relation.ChildrenID.Hex()
			ft.childrenOfChildres(&children, children.ID, queryPerson, relationsShips)
			member.Childrens = append(member.Childrens, children)
		}
	}
}

func (ft *FamilyTreeServiceImpl) updatePersonInfo(members *Members, personData []person.Person) {
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
