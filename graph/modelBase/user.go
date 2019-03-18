package modelBase

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
//将一个User根据用户ID和名字转化为一条三元组，label为nil
func (this User) Quad() quad.Quad {
	return quad.Make(this.Id, UserName.String(), this.Name, nil)
}

//将一个user添加到cayley中
func (this User) AddUserToCayley(store *cayley.Handle, ) error {
	return store.AddQuad(this.Quad())
}

//添加一个或者多个user到cayley中
func AddUserToCayley(store *cayley.Handle, users ...User) error {
	var quadSet []quad.Quad
	for _, u := range users {
		quadSet = append(quadSet, u.Quad())
	}
	return store.AddQuadSet(quadSet)
}