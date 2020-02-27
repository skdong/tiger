package hash

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := map[int]int{}
	for i, v := range nums {
		if f, ok := cache[v]; ok {
			if i-f > k {
				return true
			}
		} else {
			cache[v] = i
		}
	}
	return false
}
