package hash

import "testing"

type testAA struct {
	input  string
	desire int
}

func TestAA(t *testing.T) {
	cases := []testAA{
		testAA{input: "abcabcbb", desire: 3},
		testAA{input: " ", desire: 1},
	}
	for _, c := range cases {
		real := lengthOfLongestSubstring(c.input)
		if real != c.desire {
			t.Fatalf("case:%v real:[%d]", c, real)
		}
	}

}
