package bytedance

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	for _, interval := range intervals {
		i := 0
		for ; i < len(ans); i++ {
			if interval[1] < ans[i][0] {
				ans = append(ans[:i+1], ans[i:]...)
				ans[i] = interval
				break
			}
			if interval[0] > ans[i][1] {
				continue
			}
			if ans[i][0] > interval[0] {
				ans[i][0] = interval[0]
			}
			if ans[i][1] < interval[1] {
				ans[i][1] = interval[1]
				j := 1
				for ; i+j < len(ans) && ans[i][1] >= ans[i+j][0]; j++ {
					if ans[i][1] < ans[i+j][1] {
						ans[i][1] = ans[i+j][1]
					}
				}
				if j > 1 {
					ans = append(ans[:i+1], ans[i+j:]...)
				}
			}
			break

		}
		if len(ans) == 0 {
			ans = [][]int{interval}
		} else if i == len(ans) {
			ans = append(ans, interval)
		}
	}
	return ans
}
