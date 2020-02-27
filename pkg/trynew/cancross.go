package trynew

import "fmt"

func canCross(stones []int) bool {
	cache := map[int]bool{}
	record := map[string]bool{}
	for _, i := range stones {
		cache[i] = true
	}
	dest := stones[len(stones)-1]
	index, k := 0, 0
	queue := [][2]int{{index, k}}
	for len(queue) > 0 {
		nq := [][2]int{}
		for i := range queue {
			index, k = queue[i][0], queue[i][1]
			for d := -1; d < 2; d++ {
				step := d + k
				if step > 0 {
					if _, ok := cache[index+step]; ok {
						if index+step == dest {
							return true
						}
						if _, ok = record[fmt.Sprintf("%d,%d", index+step, step)]; ok {
							continue
						}
						record[fmt.Sprintf("%d,%d", index+step, step)] = true
						fmt.Println(index+step, step)
						nq = append(nq, [2]int{index + step, step})
					}
				}
			}

		}
		queue = nq
	}
	return false
}
