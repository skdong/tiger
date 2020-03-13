package bytedance

import "testing"

func TestFindKthLargest(t *testing.T) {
	i := []int{-1, 2, 0}
	k := 1
	t.Fatal(findKthLargest(i, k))
}
