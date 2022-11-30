package day02

import (
	"testing"

	"gotest.tools/v3/assert"
)

var TEST_PACKAGES []string = []string{"1x1x10", "2x3x4"}

func Test_parse_2x3x4_as_2_3_4(t *testing.T) {
	l, w, h := parsePresentSize("2x3x4")
	assert.Equal(t, 2, l)
	assert.Equal(t, 3, w)
	assert.Equal(t, 4, h)
}

func Test_paper_for_2_3_4_is_58(t *testing.T) {
	assert.Equal(t, 58, calculatePaperForPackage(2, 3, 4))
}

func Test_paper_for_1_1_10_is_43(t *testing.T) {
	assert.Equal(t, 43, calculatePaperForPackage(1, 1, 10))
}

func Test_total_paper_is_43_plus_58(t *testing.T) {
	assert.Equal(t, 43+58, calculateTotalPaper(TEST_PACKAGES))
}

func Test_ribbon_needed_for_2_3_4(t *testing.T) {
	result := calculateRibbonForPackage(2, 3, 4)
	assert.Equal(t, 34, result)
}

func Test_ribbon_needed_for_1_1_10(t *testing.T) {
	result := calculateRibbonForPackage(1, 1, 10)
	assert.Equal(t, 14, result)
}

func Test_total_ribbon_needed_for_both_packages(t *testing.T) {
	result := calculateTotalRibbon(TEST_PACKAGES)
	assert.Equal(t, 14+34, result)
}
