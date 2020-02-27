package leetcode

import "testing"

func TestThreeSum(t *testing.T) {
	//var input = []int{-1, 0, 1, 2, -1, -4}
	//var input = []int{}
	var input = []int{0,0,0}
	//var want = [][]int{{-1, 0, 1}, {-1, -1, 2}}
	got := threeSum(input)
	t.Errorf("%v", got)
}
