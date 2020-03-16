package dp

import "testing"

type ChangeCase struct {
	l1 []int
	l2 []int
	d  int
}

var CNCs = []ChangeCase{
	ChangeCase{
		l1: []int{5, 5, 9, 10},
		l2: []int{4, 7, 7, 13},
		d:  0},
	ChangeCase{
		l1: []int{100, 99, 98, 1, 2, 3},
		l2: []int{1, 2, 3, 4, 5, 40},
		d:  48},
}

func TestChangeNum(t *testing.T) {
	for _, cs := range CNCs {
		r := changeNumCompress(cs.l1, cs.l2)
		if cs.d != r {
			t.Fatalf("desire %d real %d", cs.d, r)
		}
	}
}
