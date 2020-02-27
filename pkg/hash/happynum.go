package hash

func isHappy(n int) bool {
	cache := map[int]int{}
	for n != 1 {
		cache[n] = 1
		target := 0
		for n != 0 {
			target += (n % 10) * (n % 10)
			n /= 10
		}
		if target == 1 {
			return true
		}
		if _, ok := cache[target]; ok {
			return false
		}
		n = target
	}
	return true
}
