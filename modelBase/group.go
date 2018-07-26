package modelBase

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
	"relation-graph/graphRelation/createTriple/conf"
)

type Group struct {
	Id int `json:"id"`
	Name string `json:"name"`
}


func (this Group) Quad() quad.Quad {
	return quad.Make(this.Id, GroupName.String(), this.Name, nil)
}

func (this Group) AddGroupToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuad(this.Quad())
}

func AddGroupToCayley(groups ...Group) error {
	var quadSet []quad.Quad
	for _, g := range groups {
		quadSet = append(quadSet, g.Quad())
	}
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(quadSet)
}