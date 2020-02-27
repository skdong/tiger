package hash

import "testing"

type FSCTestCase struct {
	A, B, C, D []int
	desire     int
}

func TestFourSumCount(t *testing.T) {
	testCases := []FSCTestCase{
		FSCTestCase{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
	}
	for _, c := range testCases {

		real := fourSumCount(c.A, c.B, c.C, c.D)
		if real != c.desire {
			t.Fatalf("%d %d", real, c.desire)
		}
	}
}
