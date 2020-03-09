package dp

/*
https://leetcode-cn.com/problems/unique-paths-ii/


*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([]int, len(obstacleGrid[0]))
	dp[0] = 1
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(obstacleGrid[0])-1]
}
