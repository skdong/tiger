package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRegress(t *testing.T){
	assert.Equal(t, true, regress("abc", "abc"))
}