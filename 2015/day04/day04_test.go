package day04

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_isHashValid_true_for_abcdef609043(t *testing.T) {
	assert.Equal(t, true, hashHas5Zeroes("abcdef609043"))
}

func Test_isHashValid_true_for_abcdef609042(t *testing.T) {
	assert.Equal(t, false, hashHas5Zeroes("abcdef609042"))
}

func Test_mine_for_abcdef_is_609043(t *testing.T) {
	assert.Equal(t, 609043, mineAdventCoin("abcdef"))
}
