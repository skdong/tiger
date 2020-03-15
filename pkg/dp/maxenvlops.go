package dp

func compare(a, b []int) bool {
	if a[0] > b[0] {
		return true
	} else if a[0] == b[0] && a[1] < b[1] {
		return true
	} else {
		return false
	}
}

func quickSort(g [][]int) {
	s, e := 0, len(g)-1
	for s < e {
		if compare(g[s], g[s+1]) {
			g[s], g[s+1] = g[s+1], g[s]
			s++
		} else if compare(g[s], g[e]) {
			g[s], g[s+1], g[e] = g[e], g[s], g[s+1]
			s++
			e--
		} else {
			e--
		}
	}
	if len(g) < 3 {
		return
	}
	quickSort(g[:s])
	quickSort(g[s+1:])
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) < 2 {
		return len(envelopes)
	}
	quickSort(envelopes)
	return lis2Branch(envelopes)
}

func lisDP(l [][]int) int {
	ans := 1
	dp := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if l[i][0] > l[j][0] &&
				l[i][1] > l[j][1] &&
				dp[j]+1 > dp[i] {

				dp[i] = dp[j] + 1
				if dp[i] > ans {
					ans = dp[i]
				}

			}
		}
	}
	return ans
}

func lis2Branch(l [][]int) int {
	ans := 0
	tail := make([]int, len(l))
	for _, e := range l {
		i, j := 0, ans
		for i < j {
			m := (i + j) >> 1
			switch {
			case e[1] < tail[m]:
				j = m
			case e[1] == tail[m]:
				i, j = m, m
			default:
				i = m + 1
			}
		}
		tail[i] = e[1]
		if i == ans {
			ans++
		}
	}
	return ans
}
