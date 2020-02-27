package hash

func firstUniqChar(s string) int {
	index := -1
	cache := map[rune]int{}
	for _, v := range s {
		if _, ok := cache[v]; ok {
			cache[v] += 1
		} else {
			cache[v] = 1
		}
	}
	for i, v := range s {
		if c, ok := cache[v]; ok {
			if c == 1 {
				return i
			}
		}
	}
	return index
}
