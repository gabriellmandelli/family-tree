package familytree

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/business/person"
	"github.com/gabriellmandelli/family-tree/internal/business/relationship"
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

func convert(input []*string) []string {
	output := make([]string, len(input))
	for i, v := range input {
		output[i] = *v
	}
	return output
}

func (ft *FamilyTreeServiceImpl) GetFamilyTree(ctx context.Context, personID string) (*FamilyTree, *errorx.Error) {
	familyTree := FamilyTree{}
	var queryPerson []*string

	relationsShips, ex := ft.relationShipService.FindAll(ctx, "", "")
	if ex != nil {
		return &familyTree, ex
	}

	ft.parentsOfParents(&familyTree.Members, personID, &queryPerson, relationsShips)
	ft.childrenOfChildres(&familyTree.Members, personID, &queryPerson, relationsShips)

	personData, ex := ft.personService.FindInBatch(ctx, convert(queryPerson))
	if ex != nil {
		return &familyTree, ex
	}

	for _, p := range personData {
		member := familyTree.Members.searchMember(p.ID.Hex())
		if member != nil {
			member.Name = p.Name
		}
	}

	return &familyTree, nil
}

func (ft *FamilyTreeServiceImpl) parentsOfParents(member *Members, memberID string, queryPerson *[]*string, relationsShips []relationship.RelationShip) {
	member.ID = memberID
	member.Parents = make([]Members, 0)

	*queryPerson = append(*queryPerson, &memberID)

	for _, relation := range relationsShips {
		if memberID == relation.ChildrenID.Hex() {
			parent := Members{}
			parent.ID = relation.ParentID.Hex()
			ft.parentsOfParents(&parent, parent.ID, queryPerson, relationsShips)
			member.Parents = append(member.Parents, parent)
		}
	}
}

func (ft *FamilyTreeServiceImpl) childrenOfChildres(member *Members, memberID string, queryPerson *[]*string, relationsShips []relationship.RelationShip) {
	member.ID = memberID
	member.Childrens = make([]Members, 0)

	*queryPerson = append(*queryPerson, &memberID)

	for _, relation := range relationsShips {

		if memberID == relation.ParentID.Hex() {
			children := Members{}
			children.ID = relation.ChildrenID.Hex()
			ft.childrenOfChildres(&children, children.ID, queryPerson, relationsShips)
			member.Childrens = append(member.Childrens, children)
		}
	}
}

func (members *Members) searchMember(id string) *Members {
	if members.ID == id {
		return members
	}
	if len(members.Parents) > 0 {
		for i := range members.Parents {
			member := members.Parents[i].searchMember(id)
			if member != nil {
				return member
			}
		}
	}
	if len(members.Childrens) > 0 {
		for i := range members.Childrens {
			member := members.Childrens[i].searchMember(id)
			if member != nil {
				return member
			}
		}
	}
	return nil
}
