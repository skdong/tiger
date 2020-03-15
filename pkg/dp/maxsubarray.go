package dp

func maxSubArray(nums []int) int {
	ans, last := nums[0], 0
	for _, i := range nums {
		if last+i > i {
			last += i
		} else {
			last = i
		}
		if last > ans {
			ans = last
		}
	}
	return ans
}
