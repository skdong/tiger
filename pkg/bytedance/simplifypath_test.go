package bytedance

import "testing"

type SPCase struct {
	s string
	d string
}

func BenchmarkSimplifyPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simplifyPath("/a//b////c/d//././/..")

	}
}

func TestSimplifyPath(t *testing.T) {
	spCases := []SPCase{
		SPCase{s: "/a//b////c/d//././/..", d: "/a/b/c"},
		SPCase{s: "/a/./b/../../c/", d: "/c"},
	}
	for _, c := range spCases {
		r := simplifyPath(c.s)
		if r != c.d {
			t.Fatal(c, r)
		}
	}
}
