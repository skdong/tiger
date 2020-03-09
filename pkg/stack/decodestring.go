package stack

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []string{}
	var attr int // 0 1 int 2 string
	var e string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if attr != 1 && len(e) > 0 {
				stack = append(stack, e)
				e = ""
			}
			attr = 1
			e = e + string(s[i])
		} else if s[i] == '[' {
			if len(e) > 0 {
				stack = append(stack, e)
				e = ""
			}
			stack = append(stack, string(s[i]))
		} else if s[i] == ']' {
			if len(e) > 0 {
				stack = append(stack, e)
				e = ""
			}
			cont := ""
			j := 0
			for j = len(stack) - 1; stack[j] != "["; j-- {
				cont = stack[j] + cont
			}
			times, _ := strconv.Atoi(stack[j-1])
			decont := ""
			for t := 0; t < times; t++ {
				decont += cont
			}
			stack[j-1] = decont
			stack = stack[:j]
		} else {
			if attr != 2 && len(e) > 0 {
				stack = append(stack, e)
				e = ""
			}
			attr = 2
			e = e + string(s[i])
			if i == len(s)-1 {
				stack = append(stack, e)
			}
		}
	}
	return strings.Join(stack, "")
}
