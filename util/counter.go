package util

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