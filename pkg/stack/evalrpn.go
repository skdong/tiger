package stack

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		curr := len(stack) - 1
		switch {
		case t == "+":
			stack[curr-1] = stack[curr-1] + stack[curr]
			stack = stack[:curr]
		case t == "-":
			stack[curr-1] = stack[curr-1] - stack[curr]
			stack = stack[:curr]
		case t == "*":
			stack[curr-1] = stack[curr-1] * stack[curr]
			stack = stack[:curr]
		case t == "/":
			stack[curr-1] = stack[curr-1] / stack[curr]
			stack = stack[:curr]
		default:
			i, _ := strconv.Atoi(t)
			stack = append(stack, i)

		}
	}
	return stack[0]
}
