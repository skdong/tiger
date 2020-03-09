package stack

func updateMatrix(matrix [][]int) [][]int {
	pathMatrix := make([][]int, len(matrix))
	stack := [][3]int{}
	for i := 0; i < len(matrix); i++ {
		pathMatrix[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				pathMatrix[i][j] = 0
				for d := -1; d < 2; d += 2 {
					if i+d < len(matrix) && i+d > -1 {
						stack = append(stack, [3]int{1, i + d, j})
					}
					if j+d < len(matrix[0]) && j+d > -1 {
						stack = append(stack, [3]int{1, i, j + d})
					}

				}
			} else {
				pathMatrix[i][j] = len(matrix) + len(matrix[0])
			}
		}
	}
	for len(stack) > 0 {
		dist, i, j := stack[0][0], stack[0][1], stack[0][2]
		stack = stack[1:]
		if pathMatrix[i][j] > dist {
			pathMatrix[i][j] = dist
			for d := -1; d < 2; d += 2 {
				if i+d < len(matrix) && i+d > -1 {
					stack = append(stack, [3]int{dist + 1, i + d, j})
				}
				if j+d < len(matrix[0]) && j+d > -1 {
					stack = append(stack, [3]int{dist + 1, i, j + d})
				}

			}
		}
	}

	return pathMatrix
}
