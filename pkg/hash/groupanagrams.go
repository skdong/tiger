package hash

func quickSort(s []byte) {
	head_i := 0
	tail_i := len(s) - 1
	for head_i < tail_i {
		if s[head_i] > s[head_i+1] {
			s[head_i], s[head_i+1] = s[head_i+1], s[head_i]
			head_i++
		} else {
			for head_i < tail_i && s[head_i] <= s[tail_i] {
				tail_i--
			}
			if s[head_i] > s[tail_i] {
				s[head_i], s[tail_i], s[head_i+1] = s[tail_i], s[head_i+1], s[head_i]
				head_i++
			}
		}
	}
	if len(s) < 3 {
		return
	}
	quickSort(s[:head_i])
	quickSort(s[head_i+1:])
}

func sortString(str string) string {
	s := []byte(str)
	quickSort(s)
	return string(s)
}
func groupAnagrams(strs []string) [][]string {
	var out [][]string
	cache := map[string][]string{}
	for _, v := range strs {
		key := sortString(v)
		if _, ok := cache[key]; ok {
			cache[key] = append(cache[key], v)
		} else {
			cache[key] = []string{v}
		}
	}
	for _, v := range cache {
		out = append(out, v)
	}
	return out
}
