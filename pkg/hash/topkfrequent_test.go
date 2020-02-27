package hash

import "testing"

type TopKTestCase struct {
	nums   []int
	k      int
	desire int
}

func TestTopKfrequent(t *testing.T) {
	input := []int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}
	k := 3
	real := topKFrequent(input, k)
	t.Fatalf("%v", real)
}
