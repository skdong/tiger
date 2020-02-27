package main

import (
	"fmt"
	"strconv"
)

func practice() {
	arrary := []int{1, 23, 4}
	for _, i := range arrary {
		fmt.Println(i)
	}

	for i, v := range "abc" {
		fmt.Println(i, string(v))
	}

	fmt.Println("abc"[0])
}

func praString() {
	str := []byte("abcd")
	if str[1] < str[2] {
		fmt.Println("hello")
	}
	fmt.Println(string(str[1:]))
	fmt.Println(string(str[:2]))
	i := 0
	for ; i < 9; i++ {
		fmt.Println(i%3*3, i/3*3)
	}
}

func main() {
	rowmaps := [9]map[byte]byte{}
	rowmaps[0] = map[byte]byte{}
	rowmaps[0][1] = 2

	fmt.Println(rowmaps)
	fmt.Println("aa" + "bb")
	fmt.Println(strconv.Itoa(1))

	tmp := "123456"
	b := []byte(tmp)
	b[2] = b[2] + 1
	fmt.Println(string(b))

	nq := [][2]int{}
	fmt.Println(len(nq))
}
