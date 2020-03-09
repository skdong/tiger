package grid

func orangesRotting(grid [][]int) int {
	count := 0
	cache := [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if x != 0 && y == 0 || x == 0 && y != 0 {
							if i+x > -1 && i+x < len(grid) && j+y > -1 && j+y < len(grid[0]) {
								if grid[i+x][j+y] == 1 {
									cache = append(cache, [2]int{i + x, j + y})
								}
							}
						}
					}
				}
			}
		}

	}

	for len(cache) > 0 {
		count++
		for _, g := range cache {
			grid[g[0]][g[1]] = 2
		}
		tCache := [][2]int{}
		for _, g := range cache {
			i, j := g[0], g[1]
			for x := -1; x < 2; x++ {
				for y := -1; y < 2; y++ {
					if x != 0 && y == 0 || x == 0 && y != 0 {
						if i+x > -1 && i+x < len(grid) && j+y > -1 && j+y < len(grid[0]) {
							if grid[i+x][j+y] == 1 {
								tCache = append(tCache, [2]int{i + x, j + y})
							}
						}
					}
				}
			}

		}
		cache = tCache

	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return count

}
