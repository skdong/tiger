package dp

func nthPersonGetsNthSeat(n int) float64 {
	stack := [][]int{make([]int, n)}
	a, r := 0.0, 0.0
	for i := 0; i < n; i++ {
		newStack := [][]int{}
		for _, site := range stack {
			if i != 0 && site[i] == 0 {
				site[i] = 1
				if i == n-1 {
					a++
					r++
				}
				newStack = append(newStack, site)
			} else {
				for j := 0; j < n; j++ {
					if site[j] == 0 {
						newSites := append([]int{}, site...)
						newSites[j] = 1
						if i == n-1 {
							a++
							if j == n-1 {
								r++
							}
						}
						newStack = append(newStack, newSites)
					}
				}
			}
		}
		stack = newStack
	}
	return r / a
}
