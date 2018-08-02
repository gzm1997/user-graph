package find

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"reflect"
	"relation-graph/graphRelation/createTriple/modelBase"
	"relation-graph/graphRelation/createTriple/util"
)



//查找跟邀请关系导致的直接关系的人
//两个返回值 类型都是map[interface{}]int
//第一个返回值是这个人邀请了谁进人他自己创建的群组
//第二个返回值是这个人被谁邀请进入他们的群组
//返回值意为每个userID对应出现在结果中多少次
func FindInviteRelevant(store *cayley.Handle, userId int) (map[int]int, map[int]int) {
	//首先假定这个人是群主
	someOneInvited := util.NewCounter()
	//这个人创建的群组里面包含的组员
	p := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.CreateGroupLink.String())).In(modelBase.ClickGroupLink.String())
	err := p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人创建的群组里面的所有人 那么这些人肯定是他邀请进来的
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			someOneInvited.Add(nativeValue.(int))
		}
	})

	//	假定这个人被邀请
	someOneInvite := util.NewCounter()
	//这个处在的群组的创建人
	p = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickGroupLink.String())).In(quad.String(modelBase.CreateGroupLink.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人处在的群组的创建人
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			someOneInvite.Add(nativeValue.(int))
		}
	})
	if err != nil {
		return nil, nil
	}
	return someOneInvited.GetInt(), someOneInvite.GetInt()
}



//具有分享文件的直接关系
func FindShareRelevant(store *cayley.Handle, userId int) (map[int]int, map[int]int) {
	//首先假定这个人分享文件给别人
	someOneShared := util.NewCounter()
	//这个人分享的所有写权限文件的接收者
	pathWriteFileLink := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.CreateWriteFileLink.String())).In(modelBase.ClickWriteFileLink.String())
	pathWriteFileLink1 := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.CreateWriteFileLink.String())).In(modelBase.ClickReadFileLink.String())
	//这个人分享的所有读权限文件的接收者
	pathReadFileLink := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.CreateReadFileLink.String())).In(modelBase.ClickReadFileLink.String())
	pathReadFileLink1 := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.CreateReadFileLink.String())).In(modelBase.ClickWriteFileLink.String())
	//查找这个人分享过文件给的人
	err := pathWriteFileLink.Or(pathWriteFileLink1).Or(pathReadFileLink).Or(pathReadFileLink1).Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人都给谁分享了读或者写权限的文件
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			someOneShared.Add(nativeValue.(int))
		}
	})
	if err != nil {
		return nil, nil
	}

	//	假定这个人被别人分享文件
	someOneShare := util.NewCounter()
	//这个人被分享的所有写权限文件的创建者
	pathWriteFileLink = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickWriteFileLink.String())).In(quad.String(modelBase.CreateWriteFileLink.String()))
	pathWriteFileLink1 = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickWriteFileLink.String())).In(quad.String(modelBase.CreateReadFileLink.String()))
	//这个人被分享的所有读权限文件的创建者
	pathReadFileLink = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickReadFileLink.String())).In(quad.String(modelBase.CreateReadFileLink.String()))
	pathReadFileLink1 = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickReadFileLink.String())).In(quad.String(modelBase.CreateWriteFileLink.String()))

	err = pathWriteFileLink.Or(pathWriteFileLink1).Or(pathReadFileLink).Or(pathReadFileLink1).Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			someOneShare.Add(nativeValue.(int))
		}
	})
	return someOneShared.GetInt(), someOneShare.GetInt()
}



//根据groupID查找这个群里面的所有成员
func FindGroupMemberByGroupId(store *cayley.Handle, groupid int) []int {
	//查询所有这个群里面所有的成员的ID
	p := cayley.StartPath(store, quad.Int(groupid)).In(quad.String(modelBase.ClickGroupLink.String()))
	var groupMember []int
	err := p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//fmt.Println("nav", nativeValue)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			groupMember = append(groupMember, nativeValue.(int))
		}
	})
	if err != nil {
		return nil
	}
	return groupMember
}


//给定一个文件ID找到有谁创建了这个文件的分享链接
func FindCreatorByFileId(store *cayley.Handle, fileId int) []int {
	pWrite := cayley.StartPath(store, quad.Int(fileId)).In(quad.String(modelBase.CreateWriteFileLink.String()))
	pRead := cayley.StartPath(store, quad.Int(fileId)).In(quad.String(modelBase.CreateReadFileLink.String()))
	var creator []int
	err := pWrite.Or(pRead).Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		creator = append(creator, nativeValue.(int))
	})
	if err != nil {
		return nil
	}
	return creator
}



//根据用户ID查找这个人创建了分享链接的文件
func FindFileIdsByUserId(store *cayley.Handle, userid int) []int {
	fileids := []int{}
	pCreateWriteFile := cayley.StartPath(store, quad.Int(userid)).Out(quad.String(modelBase.CreateWriteFileLink.String()))
	pCreateReadFile := cayley.StartPath(store, quad.Int(userid)).Out(quad.String(modelBase.CreateReadFileLink.String()))
	err := pCreateWriteFile.Or(pCreateReadFile).Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		fileids = append(fileids, nativeValue.(int))
	})
	if err != nil {
		return nil
	}
	return fileids
}


//根据文件ID查找谁点击了这个文件的分享链接
func FindClickerByFileid(store *cayley.Handle, fileid int) []int {
	clickor := []int{}
	pClickWrite := cayley.StartPath(store, quad.Int(fileid)).In(quad.String(modelBase.ClickWriteFileLink.String()))
	pClickRead := cayley.StartPath(store, quad.Int(fileid)).In(quad.String(modelBase.ClickReadFileLink.String()))
	err := pClickWrite.Or(pClickRead).Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		clickor = append(clickor, nativeValue.(int))
	})
	if err != nil {
		return nil
	}
	return clickor
}


//根据给定的groupID查找这个群里面有什么文件
func FindFileIdsByGroupId(store *cayley.Handle, groupid int) []int {
	filids := []int{}
	pFileIds := cayley.StartPath(store, quad.Int(groupid)).In(quad.String(modelBase.GroupId.String()))
	err := pFileIds.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		filids = append(filids, nativeValue.(int))
	})
	if err != nil {
		return nil
	}
	return filids
}

//根据用户ID查找他所在的群的ID
func FindGroupIdsByUserId(store *cayley.Handle, userid int) []int {
	groupids := []int{}
	p1 := cayley.StartPath(store, quad.Int(userid)).Out(quad.String(modelBase.ClickGroupLink.String()))
	p2 := cayley.StartPath(store, quad.Int(userid)).Out(quad.String(modelBase.CreateGroupLink.String()))
	err := p1.Or(p2).Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		groupids = append(groupids, nativeValue.(int))
	})
	if err != nil {
		return nil
	}
	return groupids
}