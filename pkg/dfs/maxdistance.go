package dfs

func maxDistance(grid [][]int) int {
	ans := -1
	dp := make([][]int, len(grid))
	var i, j int
	for i = 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}
	var stack [][]int
	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				stack = append(stack, []int{i, j})
			}
		}
	}
	direct := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(stack) > 0 {
		i, j = stack[0][0], stack[0][1]
		stack = stack[1:]
		for _, d := range direct {
			x, y := d[0]+i, d[1]+j
			if x >= 0 && x < len(grid) &&
				y >= 0 && y < len(grid[i]) &&
				grid[x][y] == 0 &&
				(dp[x][y] == 0 || dp[x][y] > dp[i][j]+1) {
				dp[x][y] = dp[i][j] + 1
				if dp[x][y] > ans {
					ans = dp[x][y]
				}
				stack = append(stack, []int{x, y})

			}
		}
	}

	return ans
}
