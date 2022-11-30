package lambda

import (
	"fmt"
	"testing"
)

func Test__array_FilterEle(t *testing.T) {

	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//larr := LambdaArray(arr)
	//ret1 := larr.Filter(func(ele int) bool { return ele > 5 }).Pointer().([]int)
	//fmt.Println(ret1)
	//
	//ret2 := larr.Filter(func(ele int) bool { return ele%2 == 0 }).Pointer().([]int)
	//fmt.Println(ret2)
	type user struct {
		name string
		age  int
	}
	users := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}

	ret3 := LambdaArray(users).SplitField(func(u user) string { return u.name }).Pointer().([]string)
	fmt.Println(ret3)
}
