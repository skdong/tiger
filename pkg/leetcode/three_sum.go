package leetcode


func sort(nums []int, i int, j int) {
	// sort
	k := i
	tmp := nums[k]
	left := i
	right := j - 1
	for left != right {
		if nums[right] < tmp {
			nums[k] = nums[right]
			k = right
			nums[k] = tmp
			left++
		} else if nums[left] > tmp {
			nums[k] = nums[left]
			k = left
			nums[k] = tmp
			right--
		} else if k == left {
			right--
		} else if k == right {
			left++
		}
	}
	nums[k] = tmp
	if k-i > 1 {
		sort(nums, i, k)
	}
	if j-k-1 > 1 {
		sort(nums, k+1, j)
	}

}

func threeSum(nums []int) [][]int {
	left := 0
	right := len(nums)-1
	var ret [][]int
	ret = make([][]int,0,1)
	if right < 2{
		return  ret
	}
	sort(nums,left,right+1)
	k := left + 1
	last := false
	for nums[left] <= 0 && left != right && k != right {
		sum := nums[left] + nums[k] + nums[right]
		if sum == 0 {
			if !last{
				ret = append(ret, []int{nums[left], nums[k], nums[right]})
			}
			last = true
			k++
		} else if sum > 0 {
			last = false
			right--
			k = left + 1
		} else if sum < 0 {
			last = false
			k++
			if k == right {
				left++
				k = left + 1
			}
		}
	}
	return ret
}
