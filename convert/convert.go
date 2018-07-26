package convert

import (
	"relation-graph/graphRelation/createTriple/session"
	"relation-graph/graphRelation/createTriple/model"
	"relation-graph/graphRelation/createTriple/conf"
	"relation-graph/graphRelation/createTriple/util"
)

func GetDataFromDb() ([]model.User, []model.Group, []model.CreateGroup, []model.GroupMember, []model.ShareFile) {
	session := session.GetSession()
	defer session.Close()
	var allUser []model.User
	var allGroup []model.Group

	var allCreateGroup []model.CreateGroup
	var allGroupMember []model.GroupMember
	var allShareFile []model.ShareFile

	databaseName := conf.GetDataBaseName()
	db := session.DB(databaseName)
	db.C("createGroup").Find(nil).All(&allCreateGroup)
	db.C("groupMember").Find(nil).All(&allGroupMember)
	db.C("shareFile").Find(nil).All(&allShareFile)

	//fmt.Println("createGroup", allCreateGroup)


	set := util.New()
	for _, cg := range allCreateGroup {
		set.Add(cg.User)
		set.Add(cg.Group)
	}
	for _, gm := range allGroupMember {
		set.Add(gm.User)
		set.Add(gm.Group)
	}
	for _, sf := range allShareFile {
		set.Add(sf.Subject)
		set.Add(sf.Object)
	}
	allUser, allGroup = set.GetResult()
	//model.AddManyUserToCayley(allUser)
	//model.AddManyGroupToCayley(allGroup)
	//model.AddManyCreateGroupToCayley(allCreateGroup)
	//model.AddManyGroupMemberToCayley(allGroupMember)
	//model.AddManyShareFileToCayley(allShareFile)
	return allUser, allGroup, allCreateGroup, allGroupMember, allShareFile
}