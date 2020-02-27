package trynew

import "testing"

func TestCanCross(t *testing.T) {
	//input := []int{0, 1, 3, 5, 6, 8, 12, 17}
	//input := []int{0, 1, 2, 3, 4, 8, 9, 11}
	input := []int{0, 2}
	real := canCross(input)
	t.Fatalf("%v", real)
}
