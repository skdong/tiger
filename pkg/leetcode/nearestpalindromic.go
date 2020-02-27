package leetcode

func isPalindromic(n string) bool {
	n_types := []byte(n)
	for i := 0; i < len(n_types)>>1; i++ {
		if n_types[i] != n_types[len(n_types)-1-i] {
			return false
		}
	}
	return true
}

func nearestPalindromic(n string) string {
	return n
}
