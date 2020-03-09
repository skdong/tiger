package stack

import "testing"

func TestFindTargetSum(t *testing.T) {
	input := []int{1, 1, 1, 1, 1}
	target := 3
	t.Fatal(findTargetSumWays(input, target))
}
