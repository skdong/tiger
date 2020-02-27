package hash

func countString(s []byte) int {
	var i int
	cache := map[byte]bool{}
	for i = 0; i < len(s); i++ {
		if _, ok := cache[s[i]]; ok {
			break
		} else {
			cache[s[i]] = true
		}
	}
	return i
}

func lengthOfLongestSubstring(s string) int {
	var count int
	b := []byte(s)
	for i := range b {
		c := countString(b[i:])
		if c > count {
			count = c
		}
	}

	return count
}
