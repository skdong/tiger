package huawei

func collectFiveFu(cards []string) int {
	nums := 0
	numsC := make([]int, 5)
	for _, p := range cards {
		for i := 0; i < 5; i++ {
			if p[i] == '1' {
				numsC[i]++
			}
		}
	}
	nums = numsC[0]
	for i := 1; i < 5; i++ {
		if nums > numsC[i] {
			nums = numsC[i]
		}
	}
	return nums
}
