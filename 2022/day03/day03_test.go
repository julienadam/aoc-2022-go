package day03

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_splitting_rucksacks_works_if_even(t *testing.T) {
	a, b := SplitRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Equal(t, "vJrwpWtwJgWr", a)
	assert.Equal(t, "hcsFMMfFFhFp", b)
}

func Test_find_common_item_in_compartments(t *testing.T) {
	common := FindCommonItemInCompartments("vJrwpWtwJgWr", "hcsFMMfFFhFp")
	assert.Equal(t, 'p', common)
}

func Test_priority_of_a_is_1(t *testing.T) {
	priority := getPriority('a')
	assert.Equal(t, 1, priority)
}

func Test_priority_of_z_is_26(t *testing.T) {
	priority := getPriority('z')
	assert.Equal(t, 26, priority)
}

func Test_priority_of_A_is_27(t *testing.T) {
	priority := getPriority('A')
	assert.Equal(t, 27, priority)
}

func Test_priority_of_Z_is_53(t *testing.T) {
	priority := getPriority('Z')
	assert.Equal(t, 52, priority)
}

func Test_sum_of_priorities_for_sample_is_157(t *testing.T) {
	sampleRucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	assert.Equal(t, 157, SumOfPriorities(sampleRucksacks))
}

// Part 2

func Test_can_find_badge_in_group(t *testing.T) {
	group := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg"}

	assert.Equal(t, 'r', FindBadge(group))
}

func Test_can_make_groups_of_3_rucksacks(t *testing.T) {
	sampleRucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	groups := GroupRuckSacks(sampleRucksacks)
	assert.Equal(t, 2, len(groups))
	assert.DeepEqual(t, groups[0], []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg"})
		assert.DeepEqual(t, groups[1], []string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"})
}

func Test_sum_of_priorities_for_badges_sample_is_70(t *testing.T) {
	sampleRucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	assert.Equal(t, 70, SumOfBadgePriorities(sampleRucksacks))
}
