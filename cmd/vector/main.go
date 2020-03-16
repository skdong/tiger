package main

import (
	"fmt"

	"github.com/skdong/tiger/pkg/vector"
)

func pChange(p [2][3]int) {
	p[1][2] = 90
}

func pList() {
	a := [2][3]int{{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(a)
	b := a
	fmt.Println(b)
	b[1][1] = 10
	fmt.Println(a)
	fmt.Println(b)
	pChange(b)
	fmt.Println(a)
	fmt.Println(b)
}

func pVector() {
	nums := []int{1, 2, 3}
	fmt.Println(vector.PivotIndex(nums))
}

func main() {
	pList()
}
