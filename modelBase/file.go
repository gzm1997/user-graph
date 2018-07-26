package modelBase

import (
	"github.com/cayleygraph/cayley/quad"
	"relation-graph/graphRelation/createTriple/conf"
	"github.com/cayleygraph/cayley"
)

type File struct {
	Id int
	Name string
	GroupId int
}

func (this File) Quad() []quad.Quad {
	return []quad.Quad{quad.Make(this.Id, FileName.String(), this.Name, nil), quad.Make(this.Id, GroupId.String(), this.GroupId, nil)}
}

func (this File) AddFileToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(this.Quad())
}

func AddFileToCayley(files ...File) error {
	var quadSet []quad.Quad
	for _, f := range files {
		qs := f.Quad()
		for _, q := range qs {
			quadSet = append(quadSet, q)
		}
	}
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(quadSet)
}

