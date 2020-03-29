package graph

func compare(s1, s2 string) int {
	ans := 0
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] > s2[i] {
			return 1
		} else if s1[i] < s2[i] {
			return -1
		}
	}
	if len(s1) > len(s2) {
		return 1
	} else if len(s1) < len(s2) {
		return -1
	}
	return ans
}

func appendElement(s []string, e string) []string {
	l, r := 0, len(s)
	for l < r {
		m := (l + r) >> 1
		switch compare(s[m], e) {
		case 1:
			r = m
		case -1:
			l = m + 1
		case 0:
			l, r = m, m
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

type travelState struct {
	curr       string
	ticketsMap map[string][]string
	tickets    int
	travel     []string
}

func CloneTravelState(state *travelState) *travelState {
	new := NewTravelState(state.curr, state.ticketsMap, state.tickets)
	new.travel = append(new.travel, state.travel...)
	return new
}

func NewTravelState(station string, ticketsMap map[string][]string, tickets int) *travelState {
	tmap := map[string][]string{}
	for k, v := range ticketsMap {
		nexts := append([]string{}, v...)
		tmap[k] = nexts
	}
	return &travelState{curr: station,
		ticketsMap: tmap,
		tickets:    tickets}
}

func findItinerary(tickets [][]string) []string {
	ans := []string{}
	ticketsMap := map[string][]string{}
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if toSet, ok := ticketsMap[from]; ok {
			ticketsMap[from] = appendElement(toSet, to)
		} else {
			ticketsMap[from] = []string{to}
		}
	}

	stack := []*travelState{
		NewTravelState("JFK",
			ticketsMap,
			0)}
	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		state.travel = append(state.travel, state.curr)
		if nexts, ok := state.ticketsMap[state.curr]; ok {
			for i := len(nexts) - 1; i >= 0; i-- {
				new := CloneTravelState(state)
				new.tickets++
				station := new.curr
				new.curr = nexts[i]
				if len(nexts) > 1 {
					new.ticketsMap[station] = append(new.ticketsMap[station][:i], nexts[i+1:]...)
				} else {
					delete(new.ticketsMap, station)
				}
				stack = append(stack, new)

			}
		} else if state.tickets == len(tickets) {
			return state.travel
		}

	}
	return ans
}
