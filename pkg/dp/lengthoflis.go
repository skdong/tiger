package dp

func lengthOfLIS(nums []int) int {
	ans := 0
	tail := make([]int, len(nums))
	for _, num := range nums {
		i, j := 0, ans
		for i < j {
			mid := (i + j) >> 1
			if tail[mid] < num {
				i = mid + 1
			} else {
				j = mid
			}
		}
		tail[i] = num
		if i == ans {
			ans++
		}
	}
	return ans
}
