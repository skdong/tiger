package greed

import "testing"

type rdlc struct {
	s, d string
}

var rdlcs = []rdlc{
	rdlc{
		s: "cbacdcbc",
		d: "acdb",
	},
}

func TestRemoveDuplicateLetters(t *testing.T) {
	for i, c := range rdlcs {
		r := removeDuplicateLetters(c.s)
		if r != c.d {
			t.Fatalf("%dth case err desire %v real %v", i, c.d, r)
		}
	}
}
