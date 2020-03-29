package math

import (
	"fmt"
)

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	ans := fmt.Sprintf("%d/(", nums[0])
	for i := 1; i < len(nums); i++ {
		if i == len(nums)-1 {
			ans = fmt.Sprintf("%s%d)", ans, nums[i])
		} else {
			ans = fmt.Sprintf("%s%d/", ans, nums[i])
		}

	}
	return ans
}
