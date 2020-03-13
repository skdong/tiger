package dp

func minimumTotal(triangle [][]int) int {
	var minimum int
	dp := make([]int, len(triangle[len(triangle)-1]))
	for i := 0; i < len(triangle); i++ {
		for j := i; j >= 0; j-- {
			if i > 0 && (j == i || j > 0 && dp[j-1] < dp[j]) {
				dp[j] = dp[j-1]
			}
			dp[j] += triangle[i][j]

			if i == len(triangle)-1 && (j == i || dp[j] < minimum) {
				minimum = dp[j]
			}
		}
	}
	return minimum
}
