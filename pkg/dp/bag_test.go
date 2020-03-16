package dp

import "testing"

type bagCase struct {
	v []int
	w []int
	j int
	d int
}

func initCases() []*bagCase {
	cases := []*bagCase{
		&bagCase{
			v: []int{1500, 3000, 2000},
			w: []int{1, 4, 3},
			j: 4,
			d: 3500,
		},
		&bagCase{
			w: []int{0, 2, 2, 6, 5, 4},
			v: []int{0, 6, 3, 5, 4, 6},
			j: 10,
			d: 15,
		},
	}
	return cases
}

func TestBag(t *testing.T) {
	cases := initCases()
	for _, c := range cases {

		r := bag(c.w, c.v, c.j)
		if r != c.d {
			t.Fatalf("real is %d desire is %d", r, c.d)
		}
	}
}
