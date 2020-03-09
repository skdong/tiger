package stack

import "testing"

func TestFloodFill(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2
	t.Fatal(floodFill(image, sr, sc, newColor))
}
