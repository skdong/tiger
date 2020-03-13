package dp

import "testing"

func TestMinimumTotal(t *testing.T) {
	i := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	t.Fatal(minimumTotal(i))
}
