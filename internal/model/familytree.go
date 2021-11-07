package model

type FamilyTree struct {
	Members Members `json:"members"`
}

type Members struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Parents   []Members `json:"parents"`
	Childrens []Members `json:"childrens"`
}
