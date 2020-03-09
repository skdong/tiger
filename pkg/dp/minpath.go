package dp

/*
leetcode 64.最小路径和
https://leetcode-cn.com/problems/minimum-path-sum/submissions/
*/
func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j > 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j > 0 && dp[j-1] < dp[j] {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] += grid[i][j]
			}
		}
	}
	return dp[len(grid[0])-1]
}
