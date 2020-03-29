package dp

import (
	"testing"
)

func TestNthPersonGetsNthSeat(t *testing.T) {
	i := 20
	d := 0.5
	r := nthPersonGetsNthSeat(i)
	if r != d {

		t.Fatal(d, r)
	}
}
