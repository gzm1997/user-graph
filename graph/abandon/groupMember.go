package abandon

import (
	"relation-graph/graphRelation/createTriple/session"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
	"relation-graph/graphRelation/createTriple/conf"
)

type GroupMember struct {
	Group
	User
}

func InsertAGroupMember(gm GroupMember) error {
	session := session.GetSession()
	defer session.Close()
	c := session.DB("test").C("groupMember")
	err := c.Insert(&gm)
	return err
}


func InsertManyGroupMember(gms []GroupMember) bool {
	for _, gm := range gms {
		session := session.GetSession()
		defer session.Close()
		c := session.DB("test").C("groupMember")
		err := c.Insert(&gm)
		if err != nil {
			return false
		}
	}
	return true
}

func (this GroupMember) Quad() quad.Quad {
	return quad.Make(this.Group.Id, HasMember.String(), this.User.Id, "Relation")
}

func (this GroupMember) AddGroupMemberToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuad(this.Quad())
}

func AddManyGroupMemberToCayley(groupMembers []GroupMember) error {
	var quadSet []quad.Quad
	for _, gm := range groupMembers {
		quadSet = append(quadSet, gm.Quad())
	}
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(quadSet)
}