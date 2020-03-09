package grid

import "fmt"

/*leetcode 675 为高尔夫比赛砍树*/
func cutOffTree(forest [][]int) int {
	if forest[0][0] == 0 {
		return -1
	}
	step := 0
	queue := [][2]int{}
	stack := [][2]int{{0, 0}}
	record := make([][]int, len(forest))
	for i := 0; i < len(forest); i++ {
		record[i] = make([]int, len(forest[0]))
	}
	for len(stack) > 0 {
		i, j := stack[0][0], stack[0][1]
		stack = stack[1:]
		for x := -1; x < 2; x++ {
			for y := -1; y < 2; y++ {
				if x == 0 && y != 0 || x != 0 && y == 0 {
					if i+x >= 0 && i+x < len(forest) &&
						j+y >= 0 && j+y < len(forest[0]) &&
						forest[i+x][j+y] > 0 {
						if record[i+x][j+y] == 1 {
							continue
						}
						record[i+x][j+y] = 1
						stack = append(stack, [2]int{i + x, j + y})
						if forest[i+x][j+y] > 1 {
							if len(queue) > 0 {
								qe := len(queue) - 1
								qs := 0
								for qs+1 < qe {
									bt := queue[qs+(qe-qs)/2]
									if forest[bt[0]][bt[1]] > forest[i+x][j+y] {
										qe = qs + (qe-qs)/2
									} else {
										qs = qs + (qe-qs)/2
									}

								}
								qi := qe
								if qs == 0 {
									bt := queue[qs]
									if forest[bt[0]][bt[1]] > forest[i+x][j+y] {
										qi = 0
									}
								}
								if qe == len(queue)-1 {
									bt := queue[qe]
									if forest[bt[0]][bt[1]] < forest[i+x][j+y] {
										qi = qe + 1
									}

								}
								if qi == len(queue) {

									queue = append(queue, [2]int{i + x, j + y})
								} else {

									queue = append(queue[:qi+1], queue[qi:]...)
									queue[qi] = [2]int{i + x, j + y}
								}
							} else {
								queue = [][2]int{{i + x, j + y}}
							}
						}
					}
				}
			}
		}

	}
	x, y := 0, 0
	for _, t := range queue {
		// x, y -> queuep[i][0]queue[i][1]
		stack1 := [][3]int{{0, x, y}}
		recorde1 := make([][]int, len(forest))
		for i := 0; i < len(forest); i++ {
			recorde1[i] = make([]int, len(forest[0]))
		}
		recorde1[x][y] = 1

		for len(stack1) > 0 {
			d, x, y := stack1[0][0], stack1[0][1], stack1[0][2]
			stack1 = stack1[1:]
			if x == t[0] && y == t[1] {
				step += d
				break
			}
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if i != 0 && j == 0 || i == 0 && j != 0 {
						if i+x > 0 && i+x < len(forest) &&
							j+y > 0 && j+y < len(forest) {
							if forest[i+x][j+y] > 0 && recorde1[x+i][y+j] != 1 {
								stack1 = append(stack1, [3]int{d + 1, i + x, j + y})
								recorde1[x+i][y+j] = 1
							}
						}
					}
				}
			}
		}
		x, y = t[0], t[1]
		forest[x][y] = 1
	}
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if forest[i][j] > 1 {
				return -1
			}
		}
	}
	return step
}

/*lettcode 当不能穿过未砍掉的树的情况*/
func cutOffTreeCanNotCross(forest [][]int) int {
	if forest[0][0] == 0 {
		return -1
	}
	forest[0][0] = 1

	step := 0
	index := [2]int{0, 0}
	stack := [][2]int{{0, 0}}
	heap := [][2]int{}

	record := make([][]int, len(forest))
	for i := 0; i < len(forest); i++ {
		record[i] = make([]int, len(forest[0]))
	}

	for len(stack) > 0 || len(heap) > 0 {
		for len(stack) > 0 {
			// travel all arrive tree
			i, j := stack[0][0], stack[0][1]
			stack = stack[1:]
			record[i][j] = 1
			for mv := -1; mv < 2; mv += 2 {
				if i+mv >= 0 && i+mv < len(forest) {
					if forest[i+mv][j] == 1 {
						if record[i+mv][j] != 1 {
							stack = append(stack, [2]int{i + mv, j})
							record[i+mv][j] = 1
						}
					} else if forest[i+mv][j] > 1 {
						if record[i+mv][j] != 2 {
							heap = append(heap, [2]int{i + mv, j})
							record[i+mv][j] = 2
						}
					}
				}

				if j+mv >= 0 && j+mv < len(forest[0]) {
					if forest[i][j+mv] == 1 {
						if record[i][j+mv] != 1 {
							stack = append(stack, [2]int{i, j + mv})
							record[i][j+mv] = 1
						}
					} else if forest[i][j+mv] > 1 {
						if record[i][j+mv] != 2 {
							heap = append(heap, [2]int{i, j + mv})
							record[i][j+mv] = 2
						}
					}
				}
			}
		}

		// cut tree
		if len(heap) > 0 {
			// find lowest tree in heap
			fmt.Println("heap is :")
			fmt.Println(heap)
			mi := 0
			for i := 1; i < len(heap); i++ {
				min := forest[heap[mi][0]][heap[mi][1]]
				cur := forest[heap[i][0]][heap[i][1]]
				if min > cur {
					mi = i
				}
			}
			i, j := heap[mi][0], heap[mi][1]

			// cut lowest tree, add step
			forest[i][j] = 1
			if i > index[0] {
				step += i - index[0]
			} else {
				step += index[0] - i
			}
			if j > index[1] {
				step += j - index[1]
			} else {
				step += index[1] - j
			}
			index = [2]int{i, j}
			fmt.Println(index)
			fmt.Println(step)

			// remove lowest tree from heap
			heap = append(heap[:mi], heap[mi+1:]...)

			// add cuted tree in stack
			stack = append(stack, [2]int{i, j})
		}
	}

	// check if tree not cut
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if forest[i][j] > 1 {
				return -1
			}
		}
	}
	return step
}
