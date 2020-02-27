package vector

func findDiagonalOrder(matrix [][]int) []int {
	order := make([]int, len(matrix)*len(matrix))
	for i := 0; i < len(order); i++ {
		order[i] = matrix[i][i]
	}

	return order
}
