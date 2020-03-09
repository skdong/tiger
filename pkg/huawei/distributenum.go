package huawei

func distribueNum(l1, l2 []int) int {
	diff := 0
	l := make([]int, len(l1)*2)
	min := l1[0]
	for i := 0; i < len(l1); i++ {
		l[i] = l1[i]
		l[i+len(l1)] = l2[i]
		if l1[i] < min {
			min = l1[i]
		}
		if l2[i] < min {
			min = l2[i]
		}
	}
	sum := 0
	for i := 0; i < len(l); i++ {
		l[i] -= min
		sum += l[i]
	}
	bag := sum / 2

	dp := make([]int, bag+1)
	for _, v := range l {
		for j := bag; j > 0; j-- {
			if j >= v {
				last := dp[j]
				max := v + dp[j-v]
				if last > max {
					dp[j] = max
				}
			}
		}

	}
	return diff
}
