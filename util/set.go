package util

import (
	"relation-graph/graphRelation/createTriple/model"
)

var Exists = struct {}{}

type Set struct {
	m map[interface{}]struct{}
}

func New(items ...interface{}) *Set {
	// 获取Set的地址
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

func (this Set) GetResult() ([]model.User, []model.Group) {
	//l := len(this.m)
	//resultUser := make([]model.User, l)
	//resultGroup := make([]model.Group, l)
	resultUser := []model.User{}
	resultGroup := []model.Group{}
	i := 0
	j := 0
	for k := range this.m {
		switch k.(type) {
		case model.User:
			resultUser = append(resultUser, k.(model.User))
			i++
		case model.Group:
			resultGroup = append(resultGroup, k.(model.Group))
			j++
		}
	}
	return resultUser, resultGroup
}

func (this Set) GetInt() []int {
	result := []int{}
	for k := range this.m {
		switch k.(type) {
		case int:
			result = append(result, k.(int))
		}
	}
	return result
}


func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Pop(item interface{}) interface{} {
	if s.Contains(item) {
		delete(s.m, item)
	}
	return item
}