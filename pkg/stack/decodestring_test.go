package stack

import "testing"

func TestDecodeString(t *testing.T) {
	input := "3[a2[c]]"
	t.Fatal(decodeString(input))
}
