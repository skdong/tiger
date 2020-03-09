package huawei

type Word struct {
	w string
	i []int
	m bool
}

type WT struct {
	c rune
	n []*WT
}

func minimumLengthEncodingTree(words []string) int {
	ts := []*WT{}
	tt := &ts
	for _, w := range words {
		for _, c := range w {
			co := false
			var t *WT
			for _, t = range *tt {
				if c == t.c {
					co = true
					break
				}
			}
			if !co {
				t = &WT{c: c, n: []*WT{}}
				*tt = append(*tt, t)
			}
			tt = &(t.n)
		}
	}
	return 0
}

func minimumLengthEncoding(words []string) int {
	q := make([]*Word, len(words))
	for i, word := range words {
		q[i] = &Word{w: word, i: []int{0}, m: false}
	}

	for i := 0; i < len(q)-1; i++ {
		w0 := q[i]
		if w0.m {
			continue
		}
		for _, wn := range q[i+1:] {
			if wn.m {
				continue
			}
			contain := true
			if len(wn.w) > len(w0.w) {
				for j := 0; j < len(w0.w); j++ {
					if wn.w[len(wn.w)-j-1] != w0.w[len(w0.w)-j-1] {
						contain = false
						break
					}
				}
				if contain {
					wn.i = append(wn.i, w0.i...)
					w0.m = true
					break
				}
			} else {
				for j := 0; j < len(wn.w); j++ {
					if wn.w[len(wn.w)-j-1] != w0.w[len(w0.w)-j-1] {
						contain = false
						break
					}
				}
				if contain {
					w0.i = append(w0.i, wn.i...)
					wn.m = true
				}

			}
		}
	}
	var s string
	var indexs []int
	for _, w := range q {
		if !w.m {
			for i := 0; i < len(w.i); i++ {
				indexs = append(indexs, w.i[i]+len(s))
			}
			s = s + w.w + "#"
		}
	}
	return len(s)
}
