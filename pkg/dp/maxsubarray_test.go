package dp

import "testing"

func TestMaxSubArray(t *testing.T) {
	i := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	d := 6
	r := maxSubArray(i)
	if r != d {
		t.Fatalf("desire %d, real %d", d, r)
	}
}
