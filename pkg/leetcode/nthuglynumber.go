package leetcode

func nthUglyNumber(n int, a int, b int, c int) int {
	var last int
	nth := [3]int{1, 1, 1}
	vth := [3]int{a, b, c}
	for i := 0; i < n; i++ {
		mth := 0
		for t := 0; t < 3; t++ {
			if nth[t]*vth[t] == last {
				nth[t]++
			}
			if nth[mth]*vth[mth] > nth[t]*vth[t] {
				mth = t
			}
		}
		last = nth[mth] * vth[mth]
		nth[mth]++

	}
	return last
}

func nthUglyNumberU(n int, a int, b int, c int) int {
	if a > c || a > b {
		a, b = b, a

	} else if b > c {
		b, c = c, b
	}
	return n

}
