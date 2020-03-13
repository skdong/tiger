package bytedance

func search(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		if s == e || s+1 == e {
			if target == nums[s] {
				return s
			} else if target == nums[e] {
				return e
			} else {
				return -1
			}
		}
		i := s + (e-s)/2
		if target == nums[i] {
			return i
		} else if target == nums[s] {
			return s
		} else if target == nums[e] {
			return e
		} else if nums[s] < nums[i] {
			if target > nums[i] {
				s = i
			} else if target > nums[s] {
				e = i
			} else if target < nums[e] {
				s = i
			} else {
				return -1
			}
		} else {
			if target < nums[i] {
				e = i
			} else if target < nums[e] {
				s = i
			} else if target > nums[s] {
				e = i
			} else {
				return -1
			}
		}
	}
	return -1
}
