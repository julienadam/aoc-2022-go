package day05

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_aa_has_two_consecutive_letters(t *testing.T) {
	assert.Equal(t, true, hasTwoConsecutiveEqualLetters("aa"))
}

func Test_baa_has_two_consecutive_letters(t *testing.T) {
	assert.Equal(t, true, hasTwoConsecutiveEqualLetters("baa"))
}

func Test_abc_has_not_two_consecutive_letters(t *testing.T) {
	assert.Equal(t, false, hasTwoConsecutiveEqualLetters("abc"))
}

func Test_aa_is_not_nice(t *testing.T) {
	assert.Equal(t, false, isNice("aa"))
}

func Test_xy_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("xy"))
}

func Test_axyb_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("axyb"))
}

func Test_ab_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("ab"))
}

func Test_xab_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("xab"))
}

func Test_cd_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("cd"))
}

func Test_cdx_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("cdx"))
}

func Test_pq_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("pq"))
}

func Test_ppqq_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("ppqq"))
}

func Test_ugknbfddgicrmopn_is_nice(t *testing.T) {
	assert.Equal(t, true, isNice("ugknbfddgicrmopn"))
}

func Test_aaa_is_nice(t *testing.T) {
	assert.Equal(t, true, isNice("aaa"))
}

func Test_jchzalrnumimnmhp_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("jchzalrnumimnmhp"))
}

func Test_haegwjzuvuyypxyu_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("haegwjzuvuyypxyu"))
}

func Test_dvszwmarrgswjxmb_is_naughty(t *testing.T) {
	assert.Equal(t, false, isNice("dvszwmarrgswjxmb"))
}

func Test_xyx_has_pair_with_single_char_between(t *testing.T) {
	assert.Equal(t, true, hasPairWithSingleLetterBetween("xyx"))
}

func Test_xyy_has_not_pair_with_single_char_between(t *testing.T) {
	assert.Equal(t, false, hasPairWithSingleLetterBetween("xyy"))
}

func Test_abcdefeghi_has_pair_with_single_char_between(t *testing.T) {
	assert.Equal(t, true, hasPairWithSingleLetterBetween("abcdefeghi"))
}

func Test_aaa_has_pair_with_single_char_between(t *testing.T) {
	assert.Equal(t, true, hasPairWithSingleLetterBetween("aaa"))
}

func Test_xyxy_has_repeated_non_overlapping_pair(t *testing.T) {
	assert.Equal(t, true, hasRepeatedNonOverlappingPair("xyxy"))
}

func Test_fooxybarxytru_has_repeated_non_overlapping_pair(t *testing.T) {
	assert.Equal(t, true, hasRepeatedNonOverlappingPair("fooxybarxytru"))
}

func Test_aaa_has_not_repeated_non_overlapping_pair(t *testing.T) {
	assert.Equal(t, false, hasRepeatedNonOverlappingPair("aaa"))
}

func Test_ieodomkazucvgmuy_has_not_repeated_non_overlapping_pair(t *testing.T) {
	assert.Equal(t, false, hasRepeatedNonOverlappingPair("ieodomkazucvgmuy"))
}

func Test_qjhvhtzxzqqjkmpb_is_extra_nice(t *testing.T) {
	assert.Equal(t, true, isExtraNice("qjhvhtzxzqqjkmpb"))
}

func Test_xxyxx_is_extra_nice(t *testing.T) {
	assert.Equal(t, true, isExtraNice("xxyxx"))
}

func Test_uurcxstgmygtbstg_is_not_extra_nice(t *testing.T) {
	assert.Equal(t, false, isExtraNice("uurcxstgmygtbstg"))
}

func Test_ieodomkazucvgmuy_is_not_extra_nice(t *testing.T) {
	assert.Equal(t, false, isExtraNice("ieodomkazucvgmuy"))
}
