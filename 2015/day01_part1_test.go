package day01

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_0_when_balanced(t *testing.T) {
	assert.Equal(t, 0, findFinalFloor("(())"))
}

func Test_0_when_balanced_(t *testing.T) {
	assert.Equal(t, 0, findFinalFloor("()()"))
}

func Test_3_for_3_opening(t *testing.T) {
	assert.Equal(t, 3, findFinalFloor("((("))
}

func Test_3_for_mixed(t *testing.T) {
	assert.Equal(t, 3, findFinalFloor("(()(()("))
}

func Test_3_for_mixed_inverted(t *testing.T) {
	assert.Equal(t, 3, findFinalFloor("))((((("))
}

func Test_minus_1_for_missing_open_paren(t *testing.T) {
	assert.Equal(t, -1, findFinalFloor("())"))
}

func Test_minus_1_for_missing_open_paren_inverted(t *testing.T) {
	assert.Equal(t, -1, findFinalFloor("))("))
}

func Test_minus_3_for_3_closing_parens(t *testing.T) {
	assert.Equal(t, -3, findFinalFloor(")))"))
}

func Test_minus_3_for_mixed(t *testing.T) {
	assert.Equal(t, -3, findFinalFloor(")())())"))
}
