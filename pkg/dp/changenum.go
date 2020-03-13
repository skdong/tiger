package dp

func changeNum(l1, l2 []int) int {
	ans := 0
	n := len(l1)
	l := make([]int, n*2)
	min := l1[0]
	for i := 0; i < n; i++ {
		if l1[i] < min {
			min = l1[i]
		}
		if l2[i] < min {
			min = l2[i]
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += l1[i] - min
		sum += l2[i] - min
		l[i] = l1[i] - min
		l[i+n] = l[2] - min
	}
	sum = sum / 2

	dp := make([][][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dp[i] = make([][]int, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = make([]int, n)
			if j == 0 {
				continue
			}
			for k := 0; k <= i && k < n; k++ {
				if i > 0 && k > 0 && j >= l[i] && dp[i-1][j][k] < dp[i-1][j-l[i]][k-1]+l[i] {
					dp[i][j][k] = dp[i-1][j-l[i]][k-1] + l[i]
				} else if i > 0 {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					if j >= l[i] {
						dp[i][j][k] = l[i]
					}
				}
			}
		}
	}

	return ans
}
