package find

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"reflect"
	"relation-graph/graphRelation/createTriple/model"
	"relation-graph/graphRelation/createTriple/conf"
)








func FindInviteRelevant(userId int) ([]int, []int) {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	//首先假定这个人是群主
	var someOneInvited []int
	//这个人创建的群组里面包含的组员
	p := cayley.StartPath(store, quad.Int(userId)).Out(quad.String(model.Create.String())).Out(quad.String(model.HasMember.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人创建的群组里面的所有人 除了他自己 那么这些人肯定是他邀请进来的
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			someOneInvited = append(someOneInvited, nativeValue.(int))
		}
	})

	//	假定这个人被邀请
	var someOneInvite []int
	//这个处在的群组的创建人
	p = cayley.StartPath(store, quad.Int(userId)).In(quad.String(model.HasMember.String())).In(quad.String(model.Create.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//查询这个人处在的群组的创建人 除了他自己
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int && nativeValue.(int) != userId {
			someOneInvite = append(someOneInvite, nativeValue.(int))
		}
	})
	return someOneInvited, someOneInvite
}