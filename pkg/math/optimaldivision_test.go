package math

import "testing"

func TestOptimaldivision(t *testing.T) {
	i := []int{1000, 100, 10, 2}
	t.Fatal(optimalDivision(i))
}
