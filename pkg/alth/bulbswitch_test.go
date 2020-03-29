package alth

import (
	"fmt"
	"testing"
)

func TestBulbSwitch(t *testing.T) {
	for i := 1; i < 100; i++ {
		fmt.Println(i, bulbSwitch(i))
	}
	t.Fatal("test")
}
