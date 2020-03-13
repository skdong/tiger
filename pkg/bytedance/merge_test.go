package bytedance

import "testing"

func TestMerge(t *testing.T) {
	//i := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//i := [][]int{{1, 4}, {4, 5}}
	i := [][]int{{5, 5}, {1, 2}, {2, 4}, {2, 3}, {4, 4}, {5, 5}, {2, 3}, {5, 6}, {0, 0}, {5, 6}}
	t.Fatal(merge(i))
}
