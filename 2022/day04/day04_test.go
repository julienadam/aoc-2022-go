package day04

import (
	"gotest.tools/v3/assert"
	"testing"
)

var Sample1 = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func Test_finds_range_inside_range(t *testing.T) {
	assert.Equal(t, true, isRangeCompletelyInsideOtherRange(2, 8, 3, 7))
}

func Test_finds_range_inside_range_the_other_way_around(t *testing.T) {
	assert.Equal(t, true, isRangeCompletelyInsideOtherRange(3, 7, 2, 8))
}

func Test_can_parse_ranges(t *testing.T) {
	a, b, c, d := parseRange("123-145,7889-4")
	assert.Equal(t, 123, a)
	assert.Equal(t, 145, b)
	assert.Equal(t, 7889, c)
	assert.Equal(t, 4, d)
}

func Test_the_sample_for_part1_has_two_completely_overlapping_pairs(t *testing.T) {
	assert.Equal(t, 2, findCompletelyOverlappingPairs(Sample1))
}

func Test_5_7_and_7_9_are_overlapping(t *testing.T) {
	assert.Equal(t, true, isRangeOverlappingRange(5, 7, 7, 9))
}

func Test_5_7_and_8_9_are_not_overlapping(t *testing.T) {
	assert.Equal(t, false, isRangeOverlappingRange(5, 7, 8, 9))
}

func Test_7_9_and_5_7_are_overlapping(t *testing.T) {
	assert.Equal(t, true, isRangeOverlappingRange(7, 9, 5, 9))
}

func Test_7_7_and_3_3_do_not_overlap(t *testing.T) {
	assert.Equal(t, false, isRangeOverlappingRange(7, 7, 3, 3))
}

func Test_sample2_has_4_overlapping_pairs(t *testing.T) {
	assert.Equal(t, 4, findPartiallyOverlappingPairs(Sample1))
}
