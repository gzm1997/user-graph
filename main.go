package main

import (
	"relation-graph/graphRelation/createTriple/modelBase"
	"time"
	"relation-graph/graphRelation/createTriple/convert"
	"fmt"
	"relation-graph/graphRelation/createTriple/find"
)

func generateSample() ([]modelBase.ShareFile, []modelBase.CreateGroup, []modelBase.GroupMember) {
	g1 := modelBase.Group{222, "g1"}
	g2 := modelBase.Group{223, "g2"}


	A := modelBase.User{1, "A"}
	B := modelBase.User{2, "B"}
	C := modelBase.User{3, "C"}
	D := modelBase.User{4, "D"}
	E := modelBase.User{5, "E"}
	F := modelBase.User{6, "F"}
	G := modelBase.User{7, "G"}
	H := modelBase.User{8, "H"}


	cgs := []modelBase.CreateGroup{modelBase.CreateGroup{A, g1}, modelBase.CreateGroup{H, g2}}


	gms := []modelBase.GroupMember{
		//g1
		modelBase.GroupMember{g1, A},
		modelBase.GroupMember{g1, B},
		modelBase.GroupMember{g1, C},
		//g2
		modelBase.GroupMember{g2, A},
		modelBase.GroupMember{g2, E},
		modelBase.GroupMember{g2, H},
		modelBase.GroupMember{g2, G}}

	sfs := []modelBase.ShareFile{
		//A分享给D
		modelBase.ShareFile{
			modelBase.ShareFileDesc{
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
		modelBase.ShareFile{
			ShareFileDesc: modelBase.ShareFileDesc{
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
		modelBase.ShareFile{
			modelBase.ShareFileDesc{
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
	//err := modelBase.InsertManyShareFile(sfs)
	//fmt.Println(err)
	//err = modelBase.InsertManyGroupMember(gms)
	//fmt.Println(err)
	//err = modelBase.InsertManyCreateGroup(cgs)
	//fmt.Println(err)
	var allUser []modelBase.User
	var allGroup []modelBase.Group

	var allCreateGroup []modelBase.CreateGroup
	var allGroupMember []modelBase.GroupMember
	var allShareFile []modelBase.ShareFile

	allUser, allGroup, allCreateGroup, allGroupMember, allShareFile = convert.GetDataFromDb()


	fmt.Println(modelBase.AddManyUserToCayley(allUser))
	fmt.Println(modelBase.AddManyGroupToCayley(allGroup))
	fmt.Println(modelBase.AddManyCreateGroupToCayley(allCreateGroup))
	fmt.Println(modelBase.AddManyGroupMemberToCayley(allGroupMember))
	fmt.Println(modelBase.AddManyShareFileToCayley(allShareFile))


	//查询与用户A有邀请关系的所有用户
	fmt.Println(find.FindInviteRelevant(1))

	//查询与用户A有组员隐形关系的所有用户
	fmt.Println(find.FindGroupMemberRelvent(4))

	//s := util.New()
	//s.Add(1)
	//s.Add(2)
	//fmt.Println(s.GetInt())
	//s.Pop(1)
	//s.Pop(2)
	//s.Pop(3)
	//fmt.Println(s.GetInt())
}