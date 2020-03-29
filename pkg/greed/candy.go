package greed

func candy(ratings []int) int {
	ans, n, last := 0, 0, 0
	for i := range ratings {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			n++
		} else {
			if n != 0 {
				if n > 1 {
					ans += (2 + n) * (n - 1) / 2
				}
				if last+1 > n+1 {
					ans += last + 1
				} else {
					ans += n + 1
				}
				n, last = 0, 0
			}
			ans += last + 1
			if i < len(ratings)-1 && ratings[i] < ratings[i+1] {
				last++
			} else {
				last = 0
			}
		}
	}
	return ans
}
