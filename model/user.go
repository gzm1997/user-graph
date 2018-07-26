package model

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
	"relation-graph/graphRelation/createTriple/conf"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (this User) Quad() quad.Quad {
	return quad.Make(this.Id, Name.String(), this.Name, "Name_info")
}

func (this User) AddUserToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuad(this.Quad())
}

func AddManyUserToCayley(users []User) error {
	var quadSet []quad.Quad
	for _, u := range users {
		quadSet = append(quadSet, u.Quad())
	}
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(quadSet)
}