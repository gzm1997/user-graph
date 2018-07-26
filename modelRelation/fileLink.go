package modelRelation

import (
	"github.com/cayleygraph/cayley/quad"
	"relation-graph/graphRelation/createTriple/conf"
	"github.com/cayleygraph/cayley"
	"relation-graph/graphRelation/createTriple/modelBase"
)

//type FileLink struct {
//	User
//	predicate FileLinkPredicate
//	File
//	time int
//}
//
//func (this FileLink) Quad() []quad.Quad {
//	relationQuad := quad.Make(this.User.Id, this.predicate.String(), this.File.Id, nil)
//	var time quad.Value
//	switch this.predicate {
//	case CreateWriteFileLink:
//		time = quad.String(CreateTime.String())
//	case CreateReadFileLink:
//		time = quad.String(CreateTime.String())
//	case ClickWriteFileLink:
//		time = quad.String(ClickTime.String())
//	case ClickReadFileLink:
//		time = quad.String(ClickTime.String())
//	}
//	timeQuad := quad.Make(UserId_FileId(this.User.Id, this.File.Id), time, this.time, nil)
//	return []quad.Quad{relationQuad, timeQuad}
//}
//
//func AddFileLinkToCayley(filelinks ...FileLink) error {
//	var quadSet []quad.Quad
//	for _, fl := range filelinks {
//		qs := fl.Quad()
//		for _, q := range qs {
//			quadSet = append(quadSet, q)
//		}
//	}
//	dbUrl := conf.GetDbUrl()
//	store, err := cayley.NewGraph("mongo", dbUrl, nil)
//	if err != nil {
//		panic(err)
//	}
//	return store.AddQuadSet(quadSet)
//}


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
	switch this.Permission {
	case Write:
		relationQuad = quad.Make(this.User.Id, modelBase.CreateWriteFileLink.String(), this.File.Id, nil)
	case Read:
		relationQuad = quad.Make(this.User.Id, modelBase.CreateReadFileLink.String(), this.File.Id, nil)
	}
	timeQuad := quad.Make(modelBase.UserId_FileId_Permission(this.User.Id, this.File.Id, this.Permission), modelBase.CreateTime.String(), this.CreateTime, nil)
	return []quad.Quad{relationQuad, timeQuad}
}

//将一个或者多个文件分享链接关系添加到cayley
func AddCreateFileLinkToCayley(cfls ...CreateFileLink) error {
	var quadSet []quad.Quad
	for _, cfl := range cfls {
		qs := cfl.Quad()
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
	switch this.Permission {
	case Write:
		relationQuad = quad.Make(this.User.Id, modelBase.ClickWriteFileLink.String(), this.File.Id, nil)
	case Read:
		relationQuad = quad.Make(this.User.Id, modelBase.ClickReadFileLink.String(), this.File.Id, nil)
	}
	timeQuad := quad.Make(modelBase.UserId_FileId_Permission(this.User.Id, this.File.Id, this.Permission), modelBase.ClickTime.String(), this.ClickTime, nil)
	return []quad.Quad{relationQuad, timeQuad}
}

//将一个点击文件分享链接关系添加到cayley中
func (this ClickFileLink) AddClickFileLinkToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(this.Quad())
}

//将一个或者多个点击文件分享链接关系添加到cayley中
func AddClickFileLinkToCayley(cfls ...ClickFileLink) error {
	var quadSet []quad.Quad
	for _, cfl := range cfls {
		qs := cfl.Quad()
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


