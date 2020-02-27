package stack

func isValid(s string) bool {
	stack := []rune{}
	pairs := [][2]rune{{'(', ')'}, {'{', '}'}, {'[', ']'}}
	for _, b := range s {
		for i := 0; i < 3; i++ {
			if b == pairs[i][0] {
				stack = append(stack, b)
			} else if b == pairs[i][1] {
				if len(stack) < 1 || stack[len(stack)-1] != pairs[i][0] {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}

	}
	if len(stack) > 0 {
		return false
	}
	return true
}
