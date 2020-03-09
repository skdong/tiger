package huawei

func SplitString(s string) (set []string) {
	for i := 0; i < len(s); i += 8 {
		if i+8 < len(s) {
			set = append(set, s[i:i+8])

		} else {
			tmp := []byte("00000000")
			for j := i; j < len(s); j++ {
				tmp[j-i] = s[j]
			}
			set = append(set, string(tmp))

		}
	}
	return set
}
