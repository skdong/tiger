package huawei

import (
	"errors"
	"testing"
)

type JPC struct {
	jbs []int
	i   int

	t int
	o []int
}

func (j *JPC) alert(t int, o []int) error {
	if t != j.t {
		return errors.New("time is err")
	}
	if len(o) != len(j.o) {
		return errors.New("o is err")
	}
	for i := range o {
		if o[i] != j.o[i] {

			return errors.New("o is err")
		}
	}
	return nil
}

func TestJobtoPrint(t *testing.T) {
	CS := []JPC{
		JPC{jbs: []int{1, 1, 9, 1, 1, 1},
			i: 0,
			t: 5,
			o: []int{2, 3, 4, 5, 0, 1}},
	}
	for _, jpc := range CS {
		if err := jpc.alert(JobToPrint(jpc.jbs, jpc.i)); err != nil {
			t.Fatalf("real:")
		}

	}
}
