package model

import (
	"relation-graph/graphRelation/createTriple/session"
	"github.com/cayleygraph/cayley/quad"
	"relation-graph/graphRelation/createTriple/conf"
	"github.com/cayleygraph/cayley"
)

type CreateGroup struct {
	User
	Group
}

func InsertACreateGroup(cg CreateGroup) error {
	session := session.GetSession()
	databaseName := conf.GetDataBaseName()
	defer session.Close()
	c := session.DB(databaseName).C("createGroup")
	err := c.Insert(&cg)
	return err
}


func InsertManyCreateGroup(cgs []CreateGroup) bool {
	databaseName := conf.GetDataBaseName()
	session := session.GetSession()
	defer session.Close()
	for _, cg := range cgs {
		c := session.DB(databaseName).C("createGroup")
		err := c.Insert(&cg)
		if err != nil {
			return false
		}
	}
	return true
}

func (this CreateGroup) Quad() quad.Quad {
	return quad.Make(this.User.Id, Create.String(), this.Group.Id, "Relation")
}

func (this CreateGroup) AddCreateGroupToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuad(this.Quad())
}

func AddManyCreateGroupToCayley(createGroups []CreateGroup) error {
	var quadSet []quad.Quad
	for _, cg := range createGroups {
		quadSet = append(quadSet, cg.Quad())
	}
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(quadSet)
}