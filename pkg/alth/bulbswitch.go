package alth

func bulbSwitch(n int) int {
	c := 0
	for i := 1; i <= n; i++ {
		a := 1
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				a++
			}
		}
		if a%2 == 1 {
			c++
		}

	}
	return c
}
