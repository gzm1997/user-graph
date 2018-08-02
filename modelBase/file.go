package modelBase

import (
	"github.com/cayleygraph/cayley/quad"
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

func (this File) AddFileToCayley(store *cayley.Handle) error {
	return store.AddQuadSet(this.Quad())
}

func AddFileToCayley(store *cayley.Handle, files ...File) error {
	var quadSet []quad.Quad
	for _, f := range files {
		qs := f.Quad()
		for _, q := range qs {
			quadSet = append(quadSet, q)
		}
	}
	return store.AddQuadSet(quadSet)
}

