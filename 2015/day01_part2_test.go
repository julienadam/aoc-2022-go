package day01

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_basement_at_position_1(t *testing.T) {
	assert.Equal(t, 1, findBasementInstructionIndex(")"))
}

func Test_basement_at_position_5(t *testing.T) {
	assert.Equal(t, 5, findBasementInstructionIndex("()())"))
}