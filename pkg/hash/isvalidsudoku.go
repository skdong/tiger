package hash

func checkLine(line []byte) bool {
	cache := map[byte]byte{}
	for _, v := range line {
		if v != '.' {
			if _, ok := cache[v]; ok {
				return false
			}
			cache[v] = v
		}
	}
	return true
}

func checkColumn(board [][]byte, i int) bool {
	cache := map[byte]byte{}
	for j := 0; j < len(board); j++ {
		if board[j][i] != '.' {
			if _, ok := cache[board[j][i]]; ok {
				return false
			}
			cache[board[j][i]] = board[j][i]
		}
	}
	return true
}

func checkGrid(board [][]byte, n, m int) bool {
	cache := map[byte]byte{}
	for i := n; i < n+3; i++ {
		for j := m; j < m+3; j++ {
			if board[i][j] != '.' {
				if _, ok := cache[board[i][j]]; ok {
					return false
				}
				cache[board[i][j]] = board[i][j]
			}
		}
	}
	return true
}

func isvalidsudoku(board [][]byte) bool {
	rowMaps := [9]map[byte]byte{}
	columeMaps := [9]map[byte]byte{}
	gridMaps := [9]map[byte]byte{}
	for i := 0; i < 9; i++ {
		rowMaps[i] = map[byte]byte{}
		columeMaps[i] = map[byte]byte{}
		gridMaps[i] = map[byte]byte{}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := rowMaps[i][board[i][j]]; ok {
				return false
			}
			rowMaps[i][board[i][j]] = board[i][j]

			if _, ok := columeMaps[j][board[i][j]]; ok {
				return false
			}
			columeMaps[j][board[i][j]] = board[i][j]

			if _, ok := gridMaps[i/3*3+j/3][board[i][j]]; ok {
				return false
			}
			gridMaps[i/3*3+j/3][board[i][j]] = board[i][j]
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	// check lines
	for _, line := range board {
		if !checkLine(line) {
			return false
		}
	}

	// check colume
	for c := 0; c < len(board); c++ {
		if !checkColumn(board, c) {
			return false
		}
	}

	// check grid
	for g := 0; g < 9; g++ {
		if !checkGrid(board, (g%3)*3, (g/3)*3) {
			return false
		}
	}

	return true
}
