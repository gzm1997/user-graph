package modelBase

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
)

type Group struct {
	Id int `json:"id"`
	Name string `json:"name"`
}


func (this Group) Quad() quad.Quad {
	return quad.Make(this.Id, GroupName.String(), this.Name, nil)
}

func (this Group) AddGroupToCayley(store *cayley.Handle, ) error {
	return store.AddQuad(this.Quad())
}

func AddGroupToCayley(store *cayley.Handle, groups ...Group) error {
	var quadSet []quad.Quad
	for _, g := range groups {
		quadSet = append(quadSet, g.Quad())
	}
	return store.AddQuadSet(quadSet)
}