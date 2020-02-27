package vector

func PivotIndex(nums []int) int {
	left := make([]int, len(nums)+1)
	right := make([]int, len(nums)+1)
	index := -1
	for i := 0; i < len(nums); i++ {
		left[i+1] = left[i] + nums[i]
		right[i+1] = right[i] + nums[len(nums)-i-1]
		if i >= len(nums)>>1 {
			if index == -1 && left[i+1] == right[len(nums)-i] {
				index = i
			}
			if left[len(nums)-i] == right[i+1] {
				index = len(nums) - i - 1
			}
		}
	}
	return index
}
