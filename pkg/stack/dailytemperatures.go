package stack

func dailyTemperatures(T []int) []int {
	D := make([]int, len(T))
	s := [][2]int{}
	for i, t := range T {
		for len(s) > 0 && t > s[len(s)-1][0] {
			D[s[len(s)-1][1]] = i - s[len(s)-1][1]
			s = s[:len(s)-1]
		}
		s = append(s, [2]int{t, i})
	}
	return D
}
