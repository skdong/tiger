package greed

import "testing"

type CandyCase struct {
	ratings []int
	d       int
}

var ccs = []CandyCase{
	CandyCase{
		ratings: []int{1, 2, 87, 87, 87, 2, 1},
		d:       13,
	},
	CandyCase{
		ratings: []int{1, 2, 3, 5, 4, 3, 2, 1, 4, 3, 2, 1, 3, 2, 1, 1, 2, 3, 4},
		d:       47,
	},
}

func TestCandy(t *testing.T) {
	for i, c := range ccs {
		g := candy(c.ratings)
		if g != c.d {
			t.Fatalf("%d %d %d", i, c.d, g)
		}
	}
}
