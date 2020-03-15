package dp

import "testing"

func TestLenthOfList(t *testing.T) {
	p := []int{10, 9, 2, 5, 3, 7, 101, 18}
	d := 4
	r := lengthOfLIS(p)
	if r != d {
		t.Fatalf("desire %d real %d", d, r)
	}
}
