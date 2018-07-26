package abandon

import (
	"relation-graph/graphRelation/createTriple/session"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
	"relation-graph/graphRelation/createTriple/conf"
)

type ShareFileDesc struct {
	Permission       string `json:"permission"`
	FileId         int    `json:"file_id"`
	GroupId        int    `json:"group_id"`
	LinkCreateTime int    `json:"link_create_time"`
	LinkType       string `json:"link_type"`
	Link           string `json:"link"`
	SaveTime       int    `json:"save_time"`
} 


type ShareFile struct {
	ShareFileDesc
	Subject User
	Object User
}




func InsertAShareFile(sf ShareFile) error {
	session := session.GetSession()
	defer session.Close()
	c := session.DB("test").C("shareFile")
	err := c.Insert(&sf)
	return err
}


func InsertManyShareFile(sfs []ShareFile) bool {
	for _, sf := range sfs {
		session := session.GetSession()
		defer session.Close()
		c := session.DB("test").C("shareFile")
		err := c.Insert(&sf)
		if err != nil {
			return false
		}
	}
	return true
}

func (this ShareFile) Quad() []quad.Quad {
	quadSet := []quad.Quad{quad.Make(
		this.Subject.Id, Share.String(), this.Object.Id, "Relation"),

		//每两个人之间根据唯一的文件ID确定一条分享关系
		quad.Make(quad.String(string(this.Subject.Id) + "_" + string(this.Object.Id)), this.Subject.Id, this.Object.Id, "Share_info"),

		//每一条分享关系根据一个文件ID的具体描述
		quad.Make(this.FileId, Permission.String(), this.Permission, "Desc_info"),
		quad.Make(this.FileId, GroupId.String(), this.GroupId,"Desc_info"),
		quad.Make(this.FileId, LinkCreateTime.String(), this.LinkCreateTime, "Desc_info"),
		quad.Make(this.FileId, LinkType.String(), this.LinkType, "Desc_info"),
		quad.Make(this.FileId, Link.String(), this.Link, "Desc_info"),
		quad.Make(this.FileId, SaveTime.String(), this.SaveTime, "Desc_info")}
	return quadSet
}

func (this ShareFile) AddShareFileToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(this.Quad())
}

func AddManyShareFileToCayley(sharefiles []ShareFile) error {
	var quadSet []quad.Quad
	for _, sf := range sharefiles {
		qs := sf.Quad()
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




// 查询分享文件关系的时候 存储返回值的结构体
// UserId为有分享文件关系的用户ID
// PerssionUser为每个具有分享文件关系的用户根据文件ID的分享权限
type ShareFileResult struct {
	UserId int
	PerssionUser []map[int]string
}

