package day06

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_turn_on_0_0_through_0_0_turns_on_the_first_pixel(t *testing.T) {
	grid := [1000][1000]bool{}
	turnOnThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, true, grid[0][0])
	assert.Equal(t, false, grid[0][1])
	assert.Equal(t, false, grid[1][0])
}

func Test_turn_on_0_0_through_1_1_turns_on_the_top_left_2_sided_square(t *testing.T) {
	grid := [1000][1000]bool{}
	turnOnThrough(&grid, 0, 0, 1, 1)

	assert.Equal(t, true, grid[0][0])
	assert.Equal(t, true, grid[0][1])
	assert.Equal(t, true, grid[1][0])
	assert.Equal(t, true, grid[1][1])
	assert.Equal(t, false, grid[0][2])
	assert.Equal(t, false, grid[1][2])
	assert.Equal(t, false, grid[2][0])
	assert.Equal(t, false, grid[2][1])
	assert.Equal(t, false, grid[2][2])
}

func Test_turn_off_0_0_through_0_0_turns_off_the_first_pixel(t *testing.T) {
	grid := [1000][1000]bool{}
	grid[0][0] = true
	grid[0][1] = true
	grid[1][0] = true
	turnOffThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, false, grid[0][0])
	assert.Equal(t, true, grid[0][1])
	assert.Equal(t, true, grid[1][0])
}

func Test_toggle_through_0_0_turns_on_the_first_pixel_second_toggle_turns_it_back_off(t *testing.T) {
	grid := [1000][1000]bool{}
	toggleThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, true, grid[0][0])
	assert.Equal(t, false, grid[0][1])
	assert.Equal(t, false, grid[1][0])

	toggleThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, false, grid[0][0])
	assert.Equal(t, false, grid[0][1])
	assert.Equal(t, false, grid[1][0])
}

func Test_can_parse_turn_on(t *testing.T) {
	input := "turn on 351,678 through 951,908"
	result := parseInstruction(input)
	assert.DeepEqual(t, Instruction{351, 678, 951, 908, on}, result)
}

func Test_can_parse_turn_off(t *testing.T) {
	input := "turn off 6,964 through 411,976"
	result := parseInstruction(input)
	assert.DeepEqual(t, Instruction{6, 964, 411, 976, off}, result)
}

func Test_can_parse_toggle(t *testing.T) {
	input := "toggle 393,804 through 510,976"
	result := parseInstruction(input)
	assert.DeepEqual(t, Instruction{393, 804, 510, 976, toggle}, result)
}

func Test_turn_all_on_means_a_million_lights_are_on(t *testing.T) {
	input := "turn on 0,0 through 999,999"
	instr := parseInstruction(input)
	grid := [1000][1000]bool{}
	applyInstructions(&grid, []Instruction{instr})
	assert.Equal(t, 1000000, countLightsOn(grid))
}

func Test_increase_0_0_through_0_0_add_1_to_the_top_left_pixel_brightness(t *testing.T) {
	grid := [1000][1000]int{}
	increaseThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, 1, grid[0][0])
	assert.Equal(t, 0, grid[0][1])
	assert.Equal(t, 0, grid[1][0])
}

func Test_increase_more_0_0_through_0_0_add_2_to_the_top_left_pixel_brightness(t *testing.T) {
	grid := [1000][1000]int{}
	increaseMoreThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, 2, grid[0][0])
	assert.Equal(t, 0, grid[0][1])
	assert.Equal(t, 0, grid[1][0])
}

func Test_decrease_0_0_through_0_0_does_nothing_if_top_left_pixel_is_off(t *testing.T) {
	grid := [1000][1000]int{}
	decreaseThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, 0, grid[0][0])
	assert.Equal(t, 0, grid[0][1])
	assert.Equal(t, 0, grid[1][0])
}

func Test_decrease_0_0_through_0_0_decreases_to1_if_top_left_pixel_is_2(t *testing.T) {
	grid := [1000][1000]int{}
	grid[0][0] = 2
	decreaseThrough(&grid, 0, 0, 0, 0)

	assert.Equal(t, 1, grid[0][0])
	assert.Equal(t, 0, grid[0][1])
	assert.Equal(t, 0, grid[1][0])
}

func Test_toggle_all_on_means_brightness_is_2000000(t *testing.T) {
	input := "toggle 0,0 through 999,999"
	instr := parseInstruction(input)
	grid := [1000][1000]int{}
	applyUpdatedInstructions(&grid, []Instruction{instr})
	assert.Equal(t, 2000000, countBrightness(grid))
}
