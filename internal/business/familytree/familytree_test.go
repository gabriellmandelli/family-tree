package familytree

import (
	"context"
	"reflect"
	"testing"

	"github.com/gabriellmandelli/family-tree/internal/business/person"
	"github.com/gabriellmandelli/family-tree/internal/business/relationship"
	"github.com/joomcode/errorx"
)

func TestNewFamilyTreeService(t *testing.T) {
	type args struct {
		personService       person.PersonService
		relationShipService relationship.RelationShipService
	}
	tests := []struct {
		name string
		args args
		want FamilyTreeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFamilyTreeService(tt.args.personService, tt.args.relationShipService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFamilyTreeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		input []*string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFamilyTreeServiceImpl_GetFamilyTree(t *testing.T) {
	type fields struct {
		personService       person.PersonService
		relationShipService relationship.RelationShipService
	}
	type args struct {
		ctx      context.Context
		personID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FamilyTree
		want1  *errorx.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := &FamilyTreeServiceImpl{
				personService:       tt.fields.personService,
				relationShipService: tt.fields.relationShipService,
			}
			got, got1 := ft.GetFamilyTree(tt.args.ctx, tt.args.personID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FamilyTreeServiceImpl.GetFamilyTree() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FamilyTreeServiceImpl.GetFamilyTree() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFamilyTreeServiceImpl_parentsOfParents(t *testing.T) {
	type fields struct {
		personService       person.PersonService
		relationShipService relationship.RelationShipService
	}
	type args struct {
		member         *Members
		memberID       string
		queryPerson    *[]*string
		relationsShips []relationship.RelationShip
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := &FamilyTreeServiceImpl{
				personService:       tt.fields.personService,
				relationShipService: tt.fields.relationShipService,
			}
			ft.parentsOfParents(tt.args.member, tt.args.memberID, tt.args.queryPerson, tt.args.relationsShips)
		})
	}
}

func TestFamilyTreeServiceImpl_childrenOfChildres(t *testing.T) {
	type fields struct {
		personService       person.PersonService
		relationShipService relationship.RelationShipService
	}
	type args struct {
		member         *Members
		memberID       string
		queryPerson    *[]*string
		relationsShips []relationship.RelationShip
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := &FamilyTreeServiceImpl{
				personService:       tt.fields.personService,
				relationShipService: tt.fields.relationShipService,
			}
			ft.childrenOfChildres(tt.args.member, tt.args.memberID, tt.args.queryPerson, tt.args.relationsShips)
		})
	}
}

func TestMembers_searchMember(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Parents   []Members
		Childrens []Members
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Members
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			members := &Members{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Parents:   tt.fields.Parents,
				Childrens: tt.fields.Childrens,
			}
			if got := members.searchMember(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Members.searchMember() = %v, want %v", got, tt.want)
			}
		})
	}
}
