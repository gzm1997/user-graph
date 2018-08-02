package main

import (
	"relation-graph/graphRelation/createTriple/result"
	"fmt"
)

func main() {
	//store := session.GetGraph()
	////将用户的基本信息添加到数据库中
	//A := modelBase.User{1, "A"}
	//B := modelBase.User{2, "B"}
	//C := modelBase.User{3, "C"}
	//D := modelBase.User{4, "D"}
	//E := modelBase.User{5, "E"}
	//F := modelBase.User{6, "F"}
	//G := modelBase.User{7, "G"}
	//H := modelBase.User{8, "H"}
	//users := []modelBase.User{A, B, C, D, E, F, G, H}
	//fmt.Println("add users to cayley", modelBase.AddUserToCayley(store, users...))
	//
	////将文件的基本信息添加到数据库中
	//f1 := modelBase.File{123, "fileid123", 222}
	//f2 := modelBase.File{124, "fileid124", 223}
	//f3 := modelBase.File{125, "fileid125", 223}
	//files := []modelBase.File{f1, f2, f3}
	//fmt.Println("add files to cayley", modelBase.AddFileToCayley(store, files...))
	//
	////将群组的基本信息添加到数据库中
	//g1 := modelBase.Group{222, "g1"}
	//g2 := modelBase.Group{223, "g2"}
	//groups := []modelBase.Group{g1, g2}
	//fmt.Println("add groups to cayley", modelBase.AddGroupToCayley(store, groups...))
	//
	////将两个群主邀请别人进群的关系添加到数据库中
	//
	////群主A创建了群g1 并且创建了邀请链接
	//creategl1 := modelRelation.CreateGroupShareLink{A, g1, int(time.Now().Unix())}
	////群主H创建了群g2 并且创建了邀请链接
	//creategl2 := modelRelation.CreateGroupShareLink{H, g2, int(time.Now().Unix())}
	//creategroupsharelinks := []modelRelation.CreateGroupShareLink{creategl1, creategl2}
	//fmt.Println("add creategroupsharelinks", modelRelation.AddCreateGroupShareLinkToCayley(store, creategroupsharelinks...))
	//
	//
	////将点击了群邀请分享链接的这个关系都添加到数据库中
	//clickgl1 := modelRelation.ClickGroupShareLink{B, g1, int(time.Now().Unix())}
	//clickgl2 := modelRelation.ClickGroupShareLink{C, g1, int(time.Now().Unix())}
	//
	//clickgl3 := modelRelation.ClickGroupShareLink{A, g2, int(time.Now().Unix())}
	//clickgl4 := modelRelation.ClickGroupShareLink{E, g2, int(time.Now().Unix())}
	//clickgl5 := modelRelation.ClickGroupShareLink{G, g2, int(time.Now().Unix())}
	//
	//clickgroupsharelinks := []modelRelation.ClickGroupShareLink{clickgl1, clickgl2, clickgl3, clickgl4, clickgl5}
	//fmt.Println("add clickgroupsharelinks", modelRelation.AddClickGroupShareLinkToCayley(store, clickgroupsharelinks...))
	//
	//
	//
	////将创建了文件分享链接这个关系都添加到数据库中
	//sf1 := modelRelation.CreateFileLink{A, f1, modelRelation.Write, int(time.Now().Unix())}
	//sf2 := modelRelation.CreateFileLink{H, f2, modelRelation.Read, int(time.Now().Unix())}
	//sf3 := modelRelation.CreateFileLink{H, f3, modelRelation.Write, int(time.Now().Unix())}
	//sf4 := modelRelation.CreateFileLink{H, f2, modelRelation.Read, int(time.Now().Unix())}
	//sfs := []modelRelation.CreateFileLink{sf1, sf2, sf3, sf4}
	//fmt.Println("add create file link relation to cayley", modelRelation.AddCreateFileLinkToCayley(store, sfs...))
	//
	//
	////将点击了文件分享链接这个关系都添加到数据库中
	//cf1 := modelRelation.ClickFileLink{D, f1, modelRelation.Write, int(time.Now().Unix())}
	//cf2 := modelRelation.ClickFileLink{D, f2, modelRelation.Read, int(time.Now().Unix())}
	//cf3 := modelRelation.ClickFileLink{F, f3, modelRelation.Write, int(time.Now().Unix())}
	//cf4 := modelRelation.ClickFileLink{F, f2, modelRelation.Read, int(time.Now().Unix())}
	//cfs := []modelRelation.ClickFileLink{cf1, cf2, cf3, cf4}
	//fmt.Println("add click file link relation to cayley", modelRelation.AddClickFileLinkToCayley(store, cfs...))
	//fmt.Println("")
	//
	//
	//fmt.Println("show every one relation about invite")
	//fmt.Println("-----------------------------------------------------------")
	//for _, u := range users {
	//	someOneInvitedBy, someOneInvite := find.FindInviteRelevant(store, u.Id)
	//	fmt.Println(u.Name, "invite", someOneInvitedBy)
	//	fmt.Println(someOneInvite, "invite", u.Name)
	//	fmt.Println("")
	//}
	//
	//fmt.Println("-----------------------------------------------------------")
	//
	//fmt.Println("show every one relation about share file")
	//fmt.Println("-----------------------------------------------------------")
	//for _, u := range users {
	//	createShareFileLinkTo, clickShareFileLinkFrom := find.FindShareRelevant(store, u.Id)
	//	fmt.Println(u.Name, "share file to", createShareFileLinkTo)
	//	fmt.Println(u.Name, "click file from", clickShareFileLinkFrom)
	//	fmt.Println("")
	//}
	//fmt.Println("-----------------------------------------------------------")
	//
	//fmt.Println("show every one relation about group member")
	//for _, u := range users {
	//	fmt.Println(u.Name, "has some group members", find.FindGroupMemberRelvent(store, u.Id))
	//	fmt.Println("")
	//}
	//fmt.Println("-----------------------------------------------------------")
	//
	//
	//fmt.Println("show every one relation about share the same file")
	//for _, u := range users {
	//	a, b, c, d := find.FindShareSameFileRelevant(store, u.Id)
	//	fmt.Println(u.Name, a, b, c, d)
	//	fmt.Println("")
	//}

	//fmt.Println("-----------------------------------------------------------")
	//
	//fmt.Println("show every one relation because of share file and group members")
	//for _, u := range users {
	//	r := find.FindGroupMemberBecauseOfShareFileRelevant(store, u.Id)
	//	fmt.Println(u.Name, "has a indirect relation", r)
	//}

	arr, arr1 := result.EncodeConnectorToArray(1)
	fmt.Println(arr)
	fmt.Println(arr1)
	wr := result.Calculate(arr1)
	fmt.Println(wr)
}