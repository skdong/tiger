package stack

import "testing"

func TestUpdateMatrix(t *testing.T) {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	t.Fatal(updateMatrix(input))
}
