package vector

func dominantIndex(nums []int) int {
	index := -1
	last := index
	for i := 0; i < len(nums); i++ {
		if index == -1 || nums[i] > nums[index] {
			last = index
			index = i
		}
	}
	if last == -1 || nums[index] >= nums[last]*2 {
		return index
	}
	return -1
}
