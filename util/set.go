package util

import (
	"relation-graph/graphRelation/createTriple/modelBase"
)

var Exists = struct {}{}

type Set struct {
	m map[interface{}]struct{}
}

func NewSet(items ...interface{}) *Set {
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

func (this Set) GetResult() ([]modelBase.User, []modelBase.Group, []modelBase.File) {
	//l := len(this.m)
	//resultUser := make([]modelBase.User, l)
	//resultGroup := make([]modelBase.Group, l)
	resultUser := []modelBase.User{}
	resultGroup := []modelBase.Group{}
	resultFile := []modelBase.File{}
	for k := range this.m {
		switch k.(type) {
		case modelBase.User:
			resultUser = append(resultUser, k.(modelBase.User))
		case modelBase.Group:
			resultGroup = append(resultGroup, k.(modelBase.Group))
		case modelBase.File:
			resultFile = append(resultFile, k.(modelBase.File))
		}
	}
	return resultUser, resultGroup, resultFile
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