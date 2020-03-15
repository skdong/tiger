package dp

import "testing"

func TestMinimumTotal(t *testing.T) {
	i := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	d := 11
	r := minimumTotal(i)
	if d != r {
		t.Fatalf("desirce %d , real %d", d, r)
	}
}
