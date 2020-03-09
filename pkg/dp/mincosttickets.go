package dp

func mincostTickets(days []int, cost []int) int {
	if len(cost) < 1 || len(days) < 1 {
		return 0
	}
	dp := make([]int, 365+1)
	eDay := dp[0]
	for _, d := range days {
		if d > eDay {
			eDay = d
		}
		dp[d] = 1
	}
	for d := 1; d <= eDay; d++ {
		if dp[d] == 0 {
			dp[d] = dp[d-1]
			continue
		}
		dp[d] = dp[d-1] + cost[0]

		if d >= 7 && dp[d] > dp[d-7]+cost[1] {

			dp[d] = dp[d-7] + cost[1]
		} else if d < 7 && dp[d] > cost[1] {
			dp[d] = cost[1]
		}

		if d >= 30 && dp[d] > dp[d-30]+cost[2] {
			dp[d] = dp[d-30] + cost[2]
		} else if d < 30 && dp[d] > cost[2] {
			dp[d] = cost[2]
		}
	}
	return dp[eDay]
}
