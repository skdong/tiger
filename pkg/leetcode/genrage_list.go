package leetcode

func generageArray(n int) [][]int {
	var array []int
	for i := 0; i < n; i++ {
		array = append(array, i)
	}

	var all [][]int
	var trans func(firt int)
	trans = func(first int) {
		if first == n {
			all = append(all, append([]int{}, array...))
		}
		for i := first; i < n; i++ {
			array[first], array[i] = array[i], array[first]
			trans(first + 1)
			array[first], array[i] = array[i], array[first]
		}
	}
	trans(0)
	return all
}
