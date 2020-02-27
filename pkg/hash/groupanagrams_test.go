package hash

import "testing"

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Fatalf("%v", groupAnagrams(input))
}
