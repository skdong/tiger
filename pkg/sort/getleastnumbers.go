package sort

func getLeastNumbers(arr []int, k int) []int {
	l, r := 0, len(arr)-1
	for {
		ol, or := l, r
		for l < r {
			if arr[l] > arr[l+1] {
				arr[l], arr[l+1] = arr[l+1], arr[l]
				l++
			} else if arr[l] > arr[r] {
				arr[l], arr[l+1], arr[r] = arr[r], arr[l], arr[l+1]
				l++
				r--
			} else {
				r--
			}
		}
		if l == k-1 || l == k {
			return arr[:k]
		} else if l < k-1 {
			l, r = l+1, or
		} else {
			l, r = ol, l-1
		}
	}
}
