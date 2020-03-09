package stack

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	color := image[sr][sc]
	stack := [][2]int{{sr, sc}}
	for len(stack) > 0 {
		end := len(stack) - 1
		ir, ic := stack[end][0], stack[end][1]
		stack = stack[:end]
		if image[ir][ic] != color {
			continue
		}
		image[ir][ic] = newColor
		for i := -1; i < 2; i += 2 {
			if i+ir < len(image) && i+ir > -1 && image[i+ir][ic] == color {
				stack = append(stack, [2]int{i + ir, ic})
			}
			if i+ic < len(image[0]) && i+ic > -1 && image[ir][ic+i] == color {
				stack = append(stack, [2]int{ir, ic + i})
			}
		}

	}
	return image
}
