package hash

func singleNumber(nums []int) int {
	cache := make(map[int]int)
	for _, v := range nums {
		if _, ok := cache[v]; ok {
			delete(cache, v)
		} else {
			cache[v] = v
		}
	}
	var num int
	for key, _ := range cache {
		num = key
	}
	return num
}

func singleNumberOX(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}
