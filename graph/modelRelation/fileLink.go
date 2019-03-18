package modelRelation

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
	"relation-graph/graph/modelBase"
)



type FileLinkPermission int

const (
	Write FileLinkPermission = iota
	Read
)

func (this FileLinkPermission) String() string {
	switch this {
	case Write:
		return "Write"
	case Read:
		return "Read"
	default:
		return "Unknow"
	}
}

//创建一个文件分享链接 Permission为权限
type CreateFileLink struct {
	modelBase.User
	modelBase.File
	Permission FileLinkPermission
	CreateTime int
}

//转换为三元组
func (this CreateFileLink) Quad() []quad.Quad {
	var relationQuad quad.Quad
	var timeQuad quad.Quad
	switch this.Permission {
	case Write:
		relationQuad = quad.Make(this.User.Id, modelBase.CreateWriteFileLink.String(), this.File.Id, nil)
		timeQuad = quad.Make(modelBase.UserId_FileId(this.User.Id, this.File.Id), modelBase.CreateTimeWrite.String(), this.CreateTime, nil)
	case Read:
		relationQuad = quad.Make(this.User.Id, modelBase.CreateReadFileLink.String(), this.File.Id, nil)
		timeQuad = quad.Make(modelBase.UserId_FileId(this.User.Id, this.File.Id), modelBase.CreateTimeRead.String(), this.CreateTime, nil)
	}
	return []quad.Quad{relationQuad, timeQuad}
}

//将一个文件分享链接关系添加到cayley
func (this CreateFileLink) AddCreateFileLinkToCayley(store *cayley.Handle) error {
	return store.AddQuadSet(this.Quad())
}

//将一个或者多个文件分享链接关系添加到cayley
func AddCreateFileLinkToCayley(store *cayley.Handle, cfls ...CreateFileLink) error {
	var quadSet []quad.Quad
	for _, cfl := range cfls {
		qs := cfl.Quad()
		for _, q := range qs {
			quadSet = append(quadSet, q)
		}
	}
	return store.AddQuadSet(quadSet)
}











//点击文件分享链接关系
type ClickFileLink struct {
	modelBase.User
	modelBase.File
	Permission FileLinkPermission
	ClickTime int
}

//将一个点击文件分享链接关系转为三元组
func (this ClickFileLink) Quad() []quad.Quad {
	var relationQuad quad.Quad
	var timeQuad quad.Quad
	switch this.Permission {
	case Write:
		relationQuad = quad.Make(this.User.Id, modelBase.ClickWriteFileLink.String(), this.File.Id, nil)
		timeQuad = quad.Make(modelBase.UserId_FileId(this.User.Id, this.File.Id), modelBase.ClickTimeWrite.String(), this.ClickTime, nil)
	case Read:
		relationQuad = quad.Make(this.User.Id, modelBase.ClickReadFileLink.String(), this.File.Id, nil)
		timeQuad = quad.Make(modelBase.UserId_FileId(this.User.Id, this.File.Id), modelBase.ClickTimeRead.String(), this.ClickTime, nil)
	}
	return []quad.Quad{relationQuad, timeQuad}
}

//将一个点击文件分享链接关系添加到cayley中
func (this ClickFileLink) AddClickFileLinkToCayley(store *cayley.Handle) error {
	return store.AddQuadSet(this.Quad())
}

//将一个或者多个点击文件分享链接关系添加到cayley中
func AddClickFileLinkToCayley(store *cayley.Handle, cfls ...ClickFileLink) error {
	var quadSet []quad.Quad
	for _, cfl := range cfls {
		qs := cfl.Quad()
		for _, q := range qs {
			quadSet = append(quadSet, q)
		}
	}
	return store.AddQuadSet(quadSet)
}


