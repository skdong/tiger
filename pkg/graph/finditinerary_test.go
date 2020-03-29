package graph

import "testing"

type FICase struct {
	i [][]string
	d []string
}

var cases = []FICase{
	FICase{
		i: [][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}},
		d: []string{"JFK", "ANU", "EZE", "AXA", "TIA", "ANU", "JFK", "TIA", "ANU", "TIA", "JFK"},
	},
	FICase{
		i: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
		d: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
	},
	FICase{
		i: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
		d: []string{"JFK", "NRT", "JFK", "KUL"},
	},
	FICase{
		i: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
		d: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
	},
}

func checkTravel(r, d []string) bool {
	if len(r) != len(d) {
		return false
	}
	for i := 0; i < len(r); i++ {
		if r[i] != d[i] {
			return false
		}
	}
	return true
}

func TestFindItinerary(t *testing.T) {
	for _, c := range cases {
		//r := findItinerary(c.i)
		r := findItineraryGride(c.i)
		if !checkTravel(r, c.d) {
			t.Fatalf("desire %v , real %v", c.d, r)
		}
	}
}
