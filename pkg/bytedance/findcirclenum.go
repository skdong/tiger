package bytedance

func findCircleNum(M [][]int) int {
	ans := 0
	if len(M) == 0 {
		return ans
	}
	s := make([]int, len(M))
	for i := 0; i < len(M); i++ {
		if s[i] == 1 {
			continue
		}
		ans++
		q := []int{i}
		s[i] = 1
		for len(q) > 0 {
			f := q[0]
			q := q[1:]
			for j := 0; j < len(M); j++ {
				if M[f][j] == 1 && s[j] != 1 {
					q = append(q, j)
					s[j] = 1
				}
				if M[f][j] == 1 && s[j] != 1 {
					q = append(q, j)
					s[j] = 1
				}
			}
		}
	}
	return ans
}
