package hash

import "testing"

type testCase struct {
	s      string
	t      string
	desire bool
}

func TestIsIsomorphic(t *testing.T) {
	cases := []testCase{
		testCase{s: "egg", t: "add", desire: true},
		testCase{s: "ab", t: "aa", desire: false},
		testCase{s: "aa", t: "ab", desire: false},
	}
	for _, testC := range cases {
		real := isIsomorphic(testC.s, testC.t)
		if testC.desire != real {
			t.Fatalf("%v %v", testC, real)
		}
	}
}
