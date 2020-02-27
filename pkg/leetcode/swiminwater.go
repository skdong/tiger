package leetcode

func swimInWater(grid [][]int) int {
	t := grid[0][0]
	width := len(grid)
	if t < grid[width-1][width-1] {
		t = grid[width-1][width-1]
	}
	return t
}
