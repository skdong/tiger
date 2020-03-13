package leetcode

func gdc(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func gcdOfStrings(str1, str2 string) string {

	l1, l2 := str1+str2, str2+str1
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return ""
		}
	}
	return str1[:gdc(len(str1), len(str2))]
}

func gcdOfStringsloop(str1 string, str2 string) string {
	l := gdc(len(str1), len(str2))
	for i := 0; i < l; i++ {
		if str1[i] != str2[i] {
			return ""
		}
		for j := 1; j < len(str1)/l; j++ {
			if str1[i] != str1[i+j*l] {
				return ""
			}
		}

		for j := 1; j < len(str2)/l; j++ {
			if str2[i] != str2[i+j*l] {
				return ""
			}
		}
	}
	return str1[:l]
}
