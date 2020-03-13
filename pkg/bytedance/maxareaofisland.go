package bytedance

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for a := 0; a < len(grid); a++ {
		for b := 0; b < len(grid[0]); b++ {
			if grid[a][b] == 1 {
				area := 1
				grid[a][b] = 0
				queue := [][2]int{[2]int{a, b}}
				for len(queue) > 0 {
					i, j := queue[0][0], queue[0][1]
					queue = queue[1:]
					for x := -1; x < 2; x++ {
						for y := -1; y < 2; y++ {
							if x != 0 && y == 0 || x == 0 && y != 0 {
								if i+x > -1 && i+x < len(grid) &&
									j+y > -1 && j+y < len(grid[0]) {
									if grid[i+x][j+y] == 1 {
										area++
										grid[i+x][j+y] = 0
										queue = append(queue, [2]int{i + x, j + y})
									}
								}
							}
						}
					}
				}
				if area > maxArea {
					maxArea = area
				}
			}

		}
	}
	return maxArea

}
