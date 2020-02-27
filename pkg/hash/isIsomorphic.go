package hash

func isIsomorphic(s, t string) bool {
	hashMap := map[byte]byte{}
	markMap := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if v, ok := hashMap[s[i]]; !ok {
			if _, markerd := markMap[t[i]]; markerd {
				return false
			}
			hashMap[s[i]] = t[i]
			markMap[t[i]] = s[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}
