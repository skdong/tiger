package leetcode

import (
	"fmt"
)

func Permute(nums []int) {
	fmt.Println(permute_back(nums))
}

func permute(nums []int) [][]int {
	var ret [][]int
	for _, num := range nums {
		var new_ret [][]int
		for _, line := range ret {
			for i := 0; i <= len(line); i++ {
				new_line := make([]int, 0, len(line)+1)
				new_line = append(new_line, line[:i]...)
				new_line = append(new_line, num)
				new_line = append(new_line, line[i:]...)
				new_ret = append(new_ret, new_line)
			}
		}
		if len(ret) == 0 {
			ret = [][]int{{num}}
		} else {
			ret = new_ret
		}
	}
	return ret

}

func permute_back(nums []int) [][]int {
	var ret [][]int
	n := len(nums)
	var backtrack func(int)
	backtrack = func(first int) {
		if n-1 == first {
			line := make([]int, 0, len(nums))
			line = append(line, nums...)
			ret = append(ret, line)
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	backtrack(0)
	return ret
}
