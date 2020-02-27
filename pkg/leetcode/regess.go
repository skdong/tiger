package leetcode

func regress(s, r string) bool {
	if len(s) == 0 && len(r) == 0 {
		return true
	}
	first := (s[0] == r[0] || r[0] == '.')

	if len(r) == 1 {
		if len(s) > 1 {
			return false
		}
		return first
	} else if len(r) > 1 && r[1] == '*' {
		return regress(s[:], r[2:]) || first && regress(s[1:], r[0:])
	} else {
		return first && regress(s[1:], r[1:])
	}
}
