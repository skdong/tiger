package bytedance

func findKthLargest(nums []int, k int) int {
	si, ei := 0, len(nums)-1
	sl, el := si, ei
	for {
		for si < ei {
			if nums[si] < nums[si+1] {
				nums[si+1], nums[si] = nums[si], nums[si+1]
				si++
			} else if nums[si] < nums[ei] {
				nums[si+1], nums[si], nums[ei] = nums[si], nums[ei], nums[si+1]
				si++
				ei--
			} else {
				ei--
			}
		}
		if si < k-1 {
			si, ei = si+1, el
			sl, el = si, ei
		} else if si > k-1 {
			si, ei = sl, si-1
			sl, el = si, ei
		} else {
			break
		}
	}
	return nums[k-1]
}
