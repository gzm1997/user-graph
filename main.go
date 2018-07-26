package main

import (
	"relation-graph/graphRelation/createTriple/model"
	"time"
	"relation-graph/graphRelation/createTriple/convert"
	"fmt"
	"relation-graph/graphRelation/createTriple/find"
)

func generateSample() ([]model.ShareFile, []model.CreateGroup, []model.GroupMember) {
	g1 := model.Group{222, "g1"}
	g2 := model.Group{223, "g2"}


	A := model.User{1, "A"}
	B := model.User{2, "B"}
	C := model.User{3, "C"}
	D := model.User{4, "D"}
	E := model.User{5, "E"}
	F := model.User{6, "F"}
	G := model.User{7, "G"}
	H := model.User{8, "H"}


	cgs := []model.CreateGroup{model.CreateGroup{A, g1}, model.CreateGroup{H, g2}}


	gms := []model.GroupMember{
		//g1
		model.GroupMember{g1, A},
		model.GroupMember{g1, B},
		model.GroupMember{g1, C},
		//g2
		model.GroupMember{g2, A},
		model.GroupMember{g2, E},
		model.GroupMember{g2, H},
		model.GroupMember{g2, G}}

	sfs := []model.ShareFile{
		//A分享给D
		model.ShareFile{
			model.ShareFileDesc{
				"write",
				123,
				222,
				int(time.Now().Unix()),
				"file",
				"https://pan.wps.cn/l/t173brm",
				int(time.Now().Unix())},
			A,
			D},
		// H分享给D
		model.ShareFile{
			ShareFileDesc: model.ShareFileDesc{
				"read",
				124,
				223,
				int(time.Now().Unix()),
				"file",
				"https://pan.wps.cn/l/t173brm",
				int(time.Now().Unix())},
			Subject: H,
			Object:  D},
		//	G分享给H
		model.ShareFile{
			model.ShareFileDesc{
				"write",
				125,
				223,
				int(time.Now().Unix()),
				"file",
				"https://pan.wps.cn/l/t173brm",
				int(time.Now().Unix())},
			H,
			F}}
	return sfs, cgs, gms
}


func main()  {
	//sfs, cgs, gms := generateSample()
	//err := model.InsertManyShareFile(sfs)
	//fmt.Println(err)
	//err = model.InsertManyGroupMember(gms)
	//fmt.Println(err)
	//err = model.InsertManyCreateGroup(cgs)
	//fmt.Println(err)
	var allUser []model.User
	var allGroup []model.Group

	var allCreateGroup []model.CreateGroup
	var allGroupMember []model.GroupMember
	var allShareFile []model.ShareFile

	allUser, allGroup, allCreateGroup, allGroupMember, allShareFile = convert.GetDataFromDb()


	fmt.Println(model.AddManyUserToCayley(allUser))
	fmt.Println(model.AddManyGroupToCayley(allGroup))
	fmt.Println(model.AddManyCreateGroupToCayley(allCreateGroup))
	fmt.Println(model.AddManyGroupMemberToCayley(allGroupMember))
	fmt.Println(model.AddManyShareFileToCayley(allShareFile))


	//查询与用户A有邀请关系的所有用户
	fmt.Println(find.FindInviteRelevant(1))

	//查询与用户A有组员隐形关系的所有用户
	fmt.Println(find.FindGroupMemberRelvent(1))

	//s := util.New()
	//s.Add(1)
	//s.Add(2)
	//fmt.Println(s.GetInt())
	//s.Pop(1)
	//s.Pop(2)
	//s.Pop(3)
	//fmt.Println(s.GetInt())
}