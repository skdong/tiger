package hash

func intersect(nums1, nums2 []int) []int {
	var target []int
	cache := map[int]int{}
	for _, v := range nums1 {
		if _, ok := cache[v]; !ok {
			cache[v] = 1
		} else {
			cache[v] += 1
		}

	}
	for _, v := range nums2 {
		if c, ok := cache[v]; ok {
			if c > 0 {
				target = append(target, v)
				cache[v] -= 1
			}
		}
	}
	return target
}
