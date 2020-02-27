package hash

func mergeList(A, B []int) map[int]int {
	cache := map[int]int{}
	for a := range A {
		for b := range B {
			s := A[a] + B[b]
			if _, ok := cache[s]; !ok {
				cache[s] = 1
			} else {
				cache[s] += 1
			}
		}
	}
	return cache
}
func fourSumCount(A, B, C, D []int) int {
	var count int
	AB := mergeList(A, B)
	for c := range C {
		for d := range D {
			cd := C[c] + D[d]
			if ab, ok := AB[-cd]; ok {
				count += ab
			}
		}
	}
	return count
}
