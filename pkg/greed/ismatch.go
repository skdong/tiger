package greed

import (
	"fmt"
)

var recode = map[string]bool{}

func isMatch(s, p string) bool {
	if v, ok := recode[fmt.Sprintf("%s,%s", s, p)]; ok {
		return v
	}
	ans := false
	if len(p) == 0 && len(s) == 0 {
		ans = true
	} else if len(p) > 0 && p[0] == '*' {
		j := 1
		for j < len(p) && p[j] == '*' {
			j++
		}
		if isMatch(s, p[j:]) {
			ans = true
		}

		if len(s) > 0 && isMatch(s[1:], p[j-1:]) {
			ans = true
		}
	} else if len(p) > 0 && len(s) > 0 && (p[0] == '?' || p[0] == s[0]) {
		ans = isMatch(s[1:], p[1:])
	}
	recode[fmt.Sprintf("%s,%s", s, p)] = ans
	return ans
}
