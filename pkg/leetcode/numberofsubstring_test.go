package leetcode

import "testing"

func TestNumberOfSubString(t *testing.T) {
	input := []int{1, 1, 2, 1, 1}
	t.Fatal(numberOfSubarrays(input, 3))
}
