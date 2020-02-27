package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var first, last int
	var cur_1, cur_2, cur int
	var nums []int
	all_len := len(nums1) + len(nums2)

	for i := 0; i <= all_len/2; i++ {
		if cur_1 >= len(nums1) || cur_2 >= len(nums2) {
			if cur_1 >= len(nums1) {
				cur, nums = cur_2, nums2
			} else {
				cur, nums = cur_1, nums1
			}
			if i == all_len/2 {
				first, last = last, nums[cur]
			} else {
				cur = cur - i + all_len/2
				first, last = nums[cur-1], nums[cur]
			}
			break
		}
		if nums1[cur_1] < nums2[cur_2] {
			first, last = last, nums1[cur_1]
			cur_1++
		} else {
			first, last = last, nums2[cur_2]
			cur_2++
		}
	}

	if all_len%2 == 1 {
		return float64(last)
	} else {
		return (float64(first) + float64(last)) / 2
	}
}
