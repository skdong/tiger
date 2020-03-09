package huawei

func CountOff(n, m int) []int {
	var outs []int
	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i + 1
	}
	i := 1
	num := 0
	for len(outs) < n {
		num = num % n
		if order[num] == 0 {
			num++
		} else {
			if i == m {
				outs = append(outs, order[num])
				order[num] = 0
				i = 1
				num++
			} else {
				i++
				num++
			}
		}
	}
	return outs
}
