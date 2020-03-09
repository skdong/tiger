package stack

func numIslands(grid [][]byte) int {
	num := 0
	stack := [][2]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				grid[i][j] = '0'
				stack := append(stack, [2]int{i, j})
				for len(stack) > 0 {
					cur := len(stack) - 1
					a, b := stack[cur][0], stack[cur][1]
					stack = stack[:cur]
					for c := -1; c < 2; c += 2 {
						if a+c > -1 && a+c < len(grid) && grid[a+c][b] == '1' {
							grid[a+c][b] = '0'
							stack = append(stack, [2]int{a + c, b})
						}
						if b+c > -1 && b+c < len(grid[0]) && grid[a][b+c] == '1' {
							grid[a][b+c] = '0'
							stack = append(stack, [2]int{a, b + c})
						}
					}
				}
			}
		}
	}

	return num
}
