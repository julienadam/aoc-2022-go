package day08

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_all_trees_are_visible_in_a_2_by_3_forest(t *testing.T) {
	input := [][]int{
		{3, 0},
		{2, 5},
		{4, 7},
	}

	result := findVisibleTrees(input)

	assert.Equal(t, 6, result)
}

var sample [][]int = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func Test_21_trees_are_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, 21, findVisibleTrees(sample))
}

func Test_top_left_5_is_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, true, isVisible(sample, 1, 1))
}

func Test_top_middle_5_is_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, true, isVisible(sample, 1, 2))
}

func Test_top_right_1_is_not_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, false, isVisible(sample, 1, 3))
}

func Test_left_middle_5_is_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, true, isVisible(sample, 2, 1))
}

func Test_center_3_is_not_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, false, isVisible(sample, 2, 2))
}

func Test_right_middle_3_is_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, true, isVisible(sample, 2, 3))
}

func Test_bottom_left_3_is_not_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, false, isVisible(sample, 3, 1))
}

func Test_bottom_middle_5_is_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, true, isVisible(sample, 3, 2))
}

func Test_bottom_right_4_is_not_visible_in_the_sample(t *testing.T) {
	assert.Equal(t, false, isVisible(sample, 3, 3))
}