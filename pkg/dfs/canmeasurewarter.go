package dfs

func canMeasureWater(x, y, z int) bool {
	stack := [][]int{{0, 0}}
	recode := make([][]int, x+y+1)
	for i := 0; i < x+y+1; i++ {
		recode[i] = make([]int, x+y+1)
	}
	for len(stack) > 0 {
		rx, ry := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		if rx+ry == z {
			return true
		}
		if recode[rx][ry] == 0 {
			recode[rx][ry] = 1
		} else {
			continue
		}
		stack = append(stack, []int{0, ry})
		stack = append(stack, []int{rx, 0})
		stack = append(stack, []int{x, ry})
		stack = append(stack, []int{rx, y})
		if rx > (y - ry) {
			stack = append(stack, []int{rx - (y - ry), y})
		} else {
			stack = append(stack, []int{0, ry + rx})
		}
		if ry > (x - rx) {
			stack = append(stack, []int{x, ry - (x - rx)})

		} else {
			stack = append(stack, []int{rx + ry, 0})

		}
	}
	return false
}

func canMeasureWaterM(x, y, z int) bool {
	if z == 0 {
		return true
	}
	max := x + y
	if x > y {
		x, y = y, x
	}
	for x != 0 && y%x != 0 {
		x, y = y%x, x
	}
	if x != 0 && z <= max && z%x == 0 {
		return true
	} else if x == 0 && z == max {
		return true
	}
	return false
}
