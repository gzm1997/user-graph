package model

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
	return quad.Make(this.Id, Name.String(), this.Name, "Name_info")
}

func (this Group) AddGroupToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuad(this.Quad())
}

func AddManyGroupToCayley(groups []Group) error {
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