package graph

func findElement(s []string, e string) int {
	l, r := 0, len(s)-1
	for l <= r {
		m := (l + r) >> 1
		switch compare(s[m], e) {
		case 1:
			r = m - 1
		case -1:
			l = m + 1
		case 0:
			return m
		}
	}
	return -1
}

func addElement(s []string, e string) []string {
	l, r := 0, len(s)
	for l < r {
		m := (l + r) >> 1
		switch compare(s[m], e) {
		case 1:
			r = m
		case -1:
			l = m + 1
		case 0:
			return s
		}
	}
	if l == len(s) {
		s = append(s, e)
	} else {
		s = append(s[:l+1], s[l:]...)
		s[l] = e
	}
	return s
}

type traveInfo struct {
	grid [][]int
	i    int
	n    int
	t    []int
}

func cloneInfo(info *traveInfo) *traveInfo {
	new := newInfo(info.grid, info.i)
	new.n = info.n
	new.t = append([]int{}, info.t...)
	return new
}

func newInfo(grid [][]int, i int) *traveInfo {
	newGrid := make([][]int, len(grid))
	for i := range grid {
		if grid[i] != nil {
			newGrid[i] = append([]int{}, grid[i]...)
		}
	}
	return &traveInfo{
		grid: newGrid,
		i:    i,
	}
}
func findItineraryGride(tickets [][]string) []string {
	ans := []string{}
	stations := []string{}
	for _, ticket := range tickets {
		stations = addElement(stations, ticket[0])
		stations = addElement(stations, ticket[1])
	}
	grid := make([][]int, len(stations))
	for _, ticket := range tickets {
		i, j := findElement(stations, ticket[0]), findElement(stations, ticket[1])
		if grid[i] == nil {
			grid[i] = make([]int, len(stations))
		}
		grid[i][j]++
	}
	start := findElement(stations, "JFK")
	stack := []*traveInfo{newInfo(grid, start)}
	for len(stack) > 0 {
		info := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		info.t = append(info.t, info.i)
		for i := len(stations) - 1; i >= 0; i-- {
			if info.grid[info.i] != nil && info.grid[info.i][i] > 0 {
				new := cloneInfo(info)
				stack = append(stack, new)
				new.grid[info.i][i]--
				new.i = i
				new.n++
			}
		}
		if info.n == len(tickets) {
			ans = make([]string, 0, len(info.t))
			for _, i := range info.t {
				ans = append(ans, stations[i])
			}
			return ans
		}
	}

	return ans
}
