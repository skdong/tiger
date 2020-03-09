package main

import (
	"fmt"
	"reflect"
)

//"github.com/skdong/tiger/pkg/huawei"

func test() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	a[0] = -1
	b := append(a, 1)
	fmt.Println(b)
	a[0] = -2
	c := append(a, 2)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(make([][3]int, 3))
}

func testString() {
	a := "abcderfg"
	parten := [26]int{}
	fmt.Println(a[5:])
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a[0]))
	for _, i := range a {
		parten[i-rune('a')] += 1
	}
	fmt.Println(parten)
}

func testList() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = append(a[:2+1], a[2:]...)
	fmt.Println(a)
}

func testSlic() {
	a := "abcderfg"
	fmt.Println(a[:1])
	fmt.Println(a[1:])
}

func main() {
	testSlic()
}
