package bytedance

func trap(height []int) int {
	c := 0
	s := 0
	e := len(height) - 1
	sh, eh := s, e
	sc, ec := 0, 0
	for sh <= eh {
		if height[sh] > height[s] {
			sc += height[sh] - height[s]
		} else {
			sh = s
			c += sc
			if sh == eh {
				break
			}
			sc = 0
		}
		s++
		if height[eh] > height[e] {
			ec += height[eh] - height[e]
		} else {
			c += ec
			eh = e
			ec = 0
		}
		e--
	}
	return c
}
