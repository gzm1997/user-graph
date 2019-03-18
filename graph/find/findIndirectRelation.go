package find

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"relation-graph/graph/modelBase"
	"reflect"
	"relation-graph/graph/util"
)


//查询在同一个小组的间接关系 除去组长
func FindGroupMemberRelvent(store *cayley.Handle, userId int) map[int]int {
	groupMembers := util.NewCounter()
	//这个人所在的群组的其他组员
	p := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickGroupLink.String())).In(quad.String(modelBase.ClickGroupLink.String()))
	err := p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人所在群组的所有组员 除了他自己
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			groupMembers.Add(nativeValue.(int))
		}
	})
	if err != nil {
		return nil
	}
	//查询这个人所在群组的组长 除了他自己
	p = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickGroupLink.String())).In(quad.String(modelBase.CreateGroupLink.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人所在群组的组长 除了他自己
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			groupMembers.Pop(nativeValue.(int))
		}
	})
	//查询这个人作为组长的组里面的所有组员
	p = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.CreateGroupLink.String())).In(quad.String(modelBase.ClickGroupLink.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人所在群组的所有组员 除了他自己
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			groupMembers.Pop(nativeValue.(int))
		}
	})
	if err != nil {
		return nil
	}
	return groupMembers.GetInt()
}





//同被分享同一个文件间接关系
//返回值有两个map[int][]int 分别对应写权限和读权限的情况下
//每个文件ID对应的多个具有间接关系的用户
func FindShareSameFileRelevant(store *cayley.Handle, userId int) (map[int]map[int]int, map[int]map[int]int, map[int]map[int]int, map[int]map[int]int) {
	//找到这个人点击过的写权限文件分享链接
	pathClickWriteFileLink := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickWriteFileLink.String()))
	var writeFileIds []int
	err := pathClickWriteFileLink.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人点击过的分享权限文件链接对应的文件ID 放在writeFileIds数组中
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			writeFileIds = append(writeFileIds, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, nil, nil, nil
	}

	writeWriteResult := make(map[int]map[int]int)
	writeReadResult := make(map[int]map[int]int)
	//对于上面找到的每一个文件ID
	for _, fileId := range writeFileIds {
		writeWriteCounter := util.NewCounter()
		writeReadCounter := util.NewCounter()
		//查找每个文件ID同时也被谁点击了分享链接
		//谁点了这些文件的写权限分享链接
		pathClickWriteFileLink := cayley.StartPath(store, quad.Int(fileId)).In(quad.String(modelBase.ClickWriteFileLink.String()))
		err = pathClickWriteFileLink.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
				writeWriteCounter.Add(nativeValue.(int))
			}
		})
		if err != nil {
			return nil, nil, nil, nil
		}

		//谁点了这些文件的读权限分享链接
		pathClickReadFileLink := cayley.StartPath(store, quad.Int(fileId)).In(quad.String(modelBase.ClickReadFileLink.String()))
		//谁点了这些文件的读权限分享链接
		err = pathClickReadFileLink.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
				writeReadCounter.Add(nativeValue.(int))
			}
		})
		if err != nil {
			return nil, nil, nil, nil
		}
		r1 := writeWriteCounter.GetInt()
		r2 := writeReadCounter.GetInt()
		if len(r1) != 0 {
			writeWriteResult[fileId] = r1
		}
		if len(r2) != 0 {
			writeReadResult[fileId] = r2
		}
	}




	//找到这个人点击过的读权限文件分享链接
	pathClickReadFileLink := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(modelBase.ClickReadFileLink.String()))
	var readFileIds []int
	err = pathClickReadFileLink.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人点击过的分享权限文件链接对应的文件ID 放在readFileIds数组中
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			readFileIds = append(readFileIds, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, nil, nil, nil
	}

	readReadResult := make(map[int]map[int]int)
	readWriteResult := make(map[int]map[int]int)
	//对于上面找到的每一个文件ID
	for _, fileId := range readFileIds {
		readReadCounter := util.NewCounter()
		readWriteCounter := util.NewCounter()

		//查找每个文件ID同时也被谁点击了这个分享链接
		pathClickReadFileLink = cayley.StartPath(store, quad.Int(fileId)).In(quad.String(modelBase.ClickReadFileLink.String()))
		err = pathClickReadFileLink.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
				readReadCounter.Add(nativeValue.(int))
			}
		})
		if err != nil {
			return nil, nil, nil, nil
		}

		//查找每个文件ID同时也被谁点击了这个分享链接
		pathClickWriteFileLink = cayley.StartPath(store, quad.Int(fileId)).In(quad.String(modelBase.ClickWriteFileLink.String()))
		err = pathClickWriteFileLink.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
				readWriteCounter.Add(nativeValue.(int))
			}
		})
		if err != nil {
			return nil, nil, nil, nil
		}
		r3 := readReadCounter.GetInt()
		r4 := readWriteCounter.GetInt()
		if len(r3) != 0 {
			readReadResult[fileId] = r3
		}
		if len(r4) != 0 {
			readWriteResult[fileId] = r4
		}
	}
	return writeWriteResult, writeReadResult, readReadResult, readWriteResult
}




//查找 组员与被分享了组内文件的组外人员 这个间接关系
//返回值类型是map[int][]int 意为每个相关的群组的ID对应这个群里面的成员ID
func FindGroupMemberBecauseOfShareFileRelevant(store *cayley.Handle, userid int) map[int]map[int]int {
	fileIdGroupId := make(map[int]int)
	//找出这个人被分享的所有读权限和写权限的文件ID
	pWrite := cayley.StartPath(store, quad.Int(userid)).Out(quad.String(modelBase.ClickWriteFileLink.String()))
	pRead := cayley.StartPath(store, quad.Int(userid)).Out(quad.String(modelBase.ClickReadFileLink.String()))
	//找到所有这个人被分享的所有读写权限的文件的文件ID
	err := pWrite.Or(pRead).Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		fileId := nativeValue.(int)
		fileIdGroupId[fileId], _ = FindGroupIdByFileId(store, fileId)
	})
	if err != nil {
		return nil
	}
	result := make(map[int]map[int]int)
	//对于每个文件ID 这个文件都属于一个群 找到这个群里面的所有成员 然后去除其中对这个文件有创建分享链接操作的人
	for fileId := range fileIdGroupId {
		groupMembers := util.NewCounter()
 		groupId := fileIdGroupId[fileId]
 		//找到这个文件所在群的所有成员
 		members := FindGroupMemberByGroupId(store, groupId)
		for _, u := range members {
			groupMembers.Add(u)
		}
 		//找到对这个文件有创建分享链接操作的人
 		creator := FindCreatorByFileId(store, fileId)
		for _, c := range creator {
			//然后一个个去除对这个文件有进行创建分享链接的人
			groupMembers.Pop(c)
		}
		result[groupId] = groupMembers.GetInt()
	}


	//还有一种情况是 参数中给出的userID是群里面的人 群里面的一个文件被分享给群外面的一个人 那么这个组员也跟外面那个人有这种间接关系

	//找到这个人所在的群的ID 可能有多个
	groupids := FindGroupIdsByUserId(store, userid)
	//fmt.Println("groupid", groupids)
	//找出应该被除去的一部分文件ID 这些文件被这个人创建过分享链接(避免这个人跟文件点击者存在直接的文件分享关系)
	exceptFileIds := FindFileIdsByUserId(store, userid)
	//对于这个人所在的每一个群
	for _, g := range groupids {
		fileids := util.NewSet()
		//找到这个群里面的所有文件
		fs := FindFileIdsByGroupId(store, g)
		for _, f := range fs {
			fileids.Add(f)
		}
		//去除上面那部分应该被去除的文件ID
		for _, f := range exceptFileIds {
			fileids.Pop(f)
		}
		leftFileIds := fileids.GetInt()
		clickers := util.NewCounter()
		//对于符合条件的文件 找出这个文件的点击者
		for _, f := range leftFileIds {
			cs := FindClickerByFileid(store, f)
			for _, c := range cs {
				clickers.Add(c)
			}
		}
		cr := clickers.GetInt()
		if len(cr) != 0 {
			result[g] = cr
		}
	}
	return result
}



