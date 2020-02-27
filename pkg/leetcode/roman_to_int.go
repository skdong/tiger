package leetcode

import (
	"fmt"
)

func RomanToInt(s string) int{
	return romanToInt(s)
}

func romanToInt(s string) int {
	romanMapper := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
	for _,c := range s{
		fmt.Printf("%T",c)
	}
	fmt.Println(romanMapper)
	return 0
}