package bytedance

func quickSort(nums []int) {
	si := 0
	ei := len(nums) - 1
	for si < ei {
		if nums[si] > nums[si+1] {
			nums[si], nums[si+1] = nums[si+1], nums[si]
			si++
		} else if nums[si] > nums[ei] {
			nums[si], nums[si+1], nums[ei] = nums[ei], nums[si], nums[si+1]
			ei--
			si++
		} else {
			ei--
		}
	}
	if len(nums) < 3 {
		return
	}
	quickSort(nums[:si])
	quickSort(nums[si+1:])
}

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 3 {
		return ans
	}
	quickSort(nums)
	if nums[0] > 0 {
		return ans
	}

	for l := 0; l < len(nums)-2; l++ {
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}
		m := l + 1
		r := len(nums) - 1
		for m < r {
			if nums[l]+nums[m]+nums[r] == 0 {
				ans = append(ans, []int{nums[l], nums[m], nums[r]})
				m++
				for m < r && nums[m] == nums[m-1] {
					m++
				}
				/*
					for i<len(nums)-1 && nums[i] == nums[i+1] {
						i++
					}
					break
				*/
			} else if nums[l]+nums[m]+nums[r] > 0 {
				r--
			} else {
				m++
			}
		}
	}
	return ans

}
