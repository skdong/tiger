package huawei

import "strconv"

func compress(chars []byte) int {
	write := 0
	start := 0
	for read, c := range chars {
		if read+1 == len(chars) || c != chars[read+1] {
			chars[write] = c
			write++
			if read > start {
				for _, ch := range []byte(strconv.Itoa(read - start + 1)) {
					chars[write] = ch
					write++
				}
			}
			start = read + 1
		}
	}
	return write
}
