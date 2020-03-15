package queue

var ds = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getIslandArea(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	a := 1
	grid[i][j] = 0
	queue := [][]int{[]int{i, j}}
	for len(queue) > 0 {
		i, j = queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, d := range ds {
			if i+d[0] > -1 && i+d[0] < len(grid) &&
				j+d[1] > -1 && j+d[1] < len(grid[i]) &&
				grid[i+d[0]][j+d[1]] == 1 {

				a++
				grid[i+d[0]][j+d[1]] = 0
				queue = append(queue, []int{i + d[0], j + d[1]})
			}
		}
		/*
			if i-1 > -1 && grid[i-1][j] == 1 {
				a++
				grid[i-1][j] = 0
				queue = append(queue, []int{i - 1, j})
			}
			if i+1 < len(grid) && grid[i+1][j] == 1 {
				a++
				grid[i+1][j] = 0
				queue = append(queue, []int{i - 1, j})
			}
			if j-1 > -1 && grid[i][j-1] == 1 {
				a++
				grid[i][j-1] = 0
				queue = append(queue, []int{i - 1, j})
			}
			if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
				a++
				grid[i][j+1] = 0
				queue = append(queue, []int{i - 1, j})
			}
		*/
		/*
			for x := -1; x < 2; x++ {
				for y := -1; y < 2; y++ {
					if x != 0 && y == 0 || x == 0 && y != 0 {
						if x+i < len(grid) && x+i > -1 &&
							y+j < len(grid[i]) && y+j > -1 &&
							grid[i+x][j+y] == 1 {
							a++
							grid[i+x][j+y] = 0
							queue = append(queue, []int{i + x, j + y})
						}
					}
				}

			}
		*/
	}
	return a
}

func maxAreaOfIsland(grid [][]int) int {
	var area int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			a := getIslandArea(grid, i, j)
			if a > area {
				area = a
			}
		}
	}
	return area

}
