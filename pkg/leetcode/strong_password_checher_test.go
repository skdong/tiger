package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStrongPasswordChecker(t *testing.T) {
	s := "abc123AAAc"
	e := 0
   assert.Equal(t, e, strongPasswordChecker(s))
}