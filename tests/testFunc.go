package main

import (
	"fmt"
)

func main()  {
	//c := util.NewCounter()
	//arr := []int{1, 1, 2, 3, 4, 5, 5}
	//for _, i := range arr {
	//	c.Add(i)
	//}
	//fmt.Println(arr)
	//fmt.Println(c.GetInt())
	m := make(map[interface{}]int)
	arr1 := []int{1, 1, 2, 3, 3, 4, 4, 5}
	arr2 := []int{1, 6}
	for _, i := range arr1 {
		m[i] += 1
	}
	for _, i := range arr2 {
		_, ok := m[i]
		fmt.Println("ok", ok)
		if ok {
			m[i] -= 1
		}
	}
	fmt.Println(m)
}
