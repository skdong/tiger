package dp

import "fmt"

func bag(w, v []int, j int) int {
	dp := make([][]int, len(w))
	for i := 0; i < len(w); i++ {
		dp[i] = make([]int, j+1)
	}
	for i := 0; i < len(w); i++ {
		for a := 0; a < j+1; a++ {
			if a >= w[i] {
				if i > 0 {
					up := dp[i-1][a]
					max := v[i] + dp[i-1][a-w[i]]
					if up > max {
						dp[i][a] = up
					} else {
						dp[i][a] = max
					}
				} else {
					dp[i][a] = v[i]
				}
			} else {
				if i > 0 {
					dp[i][a] = dp[i-1][a]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[len(w)-1][j]
}

func compressBag(w, v []int, j int) int {
	dp := make([]int, j+1)
	for i := 0; i < len(w); i++ {
		for b := j; b >= w[i]; b-- {
			if b >= w[i] {
				last := dp[b]
				curmax := v[i] + dp[b-w[i]]
				if curmax > last {
					dp[b] = curmax
				}
			}
			fmt.Println(dp)
		}

	}
	return dp[j]
}
