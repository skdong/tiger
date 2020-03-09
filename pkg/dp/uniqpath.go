package dp

/*
leetcode 62.不通路径
https://leetcode-cn.com/problems/unique-paths/

到(i,j)节点的不通路径等于(i-1,j)与(i,j-1)的和

dp 记录第i行对应的不通路径

*/
func uniquePaths(m int, n int) int {
	dp := make([]int, m)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[m-1]

}
