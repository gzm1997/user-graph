package util

import (
	"relation-graph/graph/modelBase"
)

type Counter struct {
	m map[interface{}]int
}

func NewCounter(items ...interface{}) *Counter {
	c := &Counter{}
	c.m = make(map[interface{}]int)
	c.Add(items...)
	return c
}

func (this *Counter) Add(items ...interface{}) error {
	for _, item := range items {
		this.m[item] += 1
	}
	return nil
}

func (this Counter) Contain(item interface{}) bool {
	_, ok := this.m[item]
	return ok
}

func (this *Counter) Pop(item interface{}) interface{} {
	if this.Contain(item) {
		delete(this.m, item)
	}
	return item
}


func (this Counter) GetResult() map[interface{}]int {
	return this.m
}

func (this Counter) GetInt() map[int]int {
	result := make(map[int]int)
	for k := range this.m {
		switch k.(type) {
		case int:
			result[k.(int)] = this.m[k]
		}
	}
	return result
}

func GetMayKnow(userids []int, power []float64) []int {
	max := power[0]
	mayKnow := NewSet()
	mayKnow.Add(userids[0])
	for i := 1; i < len(power); i++ {
		if power[i] == max {
			//fmt.Println("等于最大值", userids[i])
			mayKnow.Add(userids[i])
		} else if power[i] > max {
			//fmt.Println("大于最大值", userids[i])
			max = power[i]
			mayKnow = NewSet()
			mayKnow.Add(userids[i])
		}
	}
	return mayKnow.GetInt()
}

func LocateUser(userids []int, users []modelBase.User) []modelBase.User {
	var result []modelBase.User
	for i := 0; i < len(userids) ; i++ {
		for j := 0; j < len(users); j++ {
			if users[j].Id == userids[i] {
				result = append(result, users[j])
			}
		}
	}
	return result
}
