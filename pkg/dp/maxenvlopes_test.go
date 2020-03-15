package dp

import "testing"

type TCEnvelope struct {
	i [][]int
	d int
}

func TestMaxEnvlopes(t *testing.T) {
	TCs := []TCEnvelope{
		TCEnvelope{
			i: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			d: 3,
		},
		TCEnvelope{
			i: [][]int{{1, 1}, {1, 1}, {1, 1}},
			d: 1,
		},
	}
	for _, cs := range TCs {
		g := maxEnvelopes(cs.i)
		if g != cs.d {
			t.Fatalf("desair %d get %d", cs.d, g)
		}
	}
}

func BenchmarkMaxEnvlopes(b *testing.B) {
	c := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	for i := 0; i < b.N; i++ {
		maxEnvelopes(c)
	}
}
