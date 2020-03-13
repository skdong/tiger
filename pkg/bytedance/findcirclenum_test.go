package bytedance

import "testing"

func TestFindCircleNum(t *testing.T) {
	i := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	t.Fatal(findCircleNum(i))
}
