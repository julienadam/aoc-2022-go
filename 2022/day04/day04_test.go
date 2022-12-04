package day04

import (
	"github.com/julienadam/adventofcode2022/util"
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

	assert.Equal(t, true, util.MakeRange(2, 8).ContainsOrIsContainedIn(util.MakeRange(3, 7)))
}

func Test_finds_range_inside_range_the_other_way_around(t *testing.T) {
	assert.Equal(t, true, util.MakeRange(3, 7).ContainsOrIsContainedIn(util.MakeRange(2, 8)))
}

func Test_can_parse_ranges2(t *testing.T) {
	r1, r2 := parseRanges("123-145,7889-4")
	assert.Equal(t, 123, r1.Start)
	assert.Equal(t, 145, r1.End)
	assert.Equal(t, 7889, r2.Start)
	assert.Equal(t, 4, r2.End)
}

func Test_the_sample_for_part1_has_two_completely_overlapping_pairs(t *testing.T) {
	assert.Equal(t, 2, findCompletelyOverlappingPairs(Sample1))
}

func Test_5_7_and_7_9_are_overlapping(t *testing.T) {
	assert.Equal(t, true, util.MakeRange(5, 7).Overlaps(util.MakeRange(7, 9)))
}

func Test_5_7_and_8_9_are_not_overlapping(t *testing.T) {
	assert.Equal(t, false, util.MakeRange(5, 7).Overlaps(util.MakeRange(8, 9)))
}

func Test_7_9_and_5_7_are_overlapping(t *testing.T) {
	assert.Equal(t, true, util.MakeRange(7, 9).Overlaps(util.MakeRange(5, 7)))
}

func Test_7_7_and_3_3_do_not_overlap(t *testing.T) {
	assert.Equal(t, false, util.MakeRange(7, 7).Overlaps(util.MakeRange(3, 3)))
}

func Test_sample2_has_4_overlapping_pairs(t *testing.T) {
	assert.Equal(t, 4, findPartiallyOverlappingPairs(Sample1))
}
