package find

import (
	"relation-graph/graphRelation/createTriple/conf"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"relation-graph/graphRelation/createTriple/model"
	"reflect"
	"relation-graph/graphRelation/createTriple/util"
)

func FindGroupMemberRelvent(userId int) []int {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	groupMembers := util.New()
	//这个人所在的群组的其他组员
	p := cayley.StartPath(store, quad.Int(userId)).In(quad.String(model.HasMember.String())).Out(quad.String(model.HasMember.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人所在群组的所有组员 除了他自己
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			groupMembers.Add(nativeValue.(int))
		}
	})
	p = cayley.StartPath(store, quad.Int(userId)).In(quad.String(model.HasMember.String())).In(quad.String(model.Create.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人所在群组的组长 除了他自己
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			groupMembers.Pop(nativeValue.(int))
		}
	})
	//查询这个人作为组长的组里面的所有组员
	p = cayley.StartPath(store, quad.Int(userId)).Out(quad.String(model.Create.String())).Out(quad.String(model.HasMember.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人所在群组的所有组员 除了他自己
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			groupMembers.Pop(nativeValue.(int))
		}
	})
	return groupMembers.GetInt()
}