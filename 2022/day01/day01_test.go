package day01

import (
	"path/filepath"
	"testing"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"gotest.tools/v3/assert"
)

const simpleSampleInput = `1000\r\n\r\n2000\r\n3000`

func Test_splitSingleElfItems(t *testing.T) {
	result := splitSingleElfItems("1000\r\n2000")
	assert.DeepEqual(t, []int{1000, 2000}, result)
}

func Test_splitSingleElfItemsWorkWithWhitespace(t *testing.T) {
	result := splitSingleElfItems("\t\t  1000\r\n2000  \t")
	assert.DeepEqual(t, []int{1000, 2000}, result)
}

func Test_splitting_input(t *testing.T) {
	split := splitByElf(simpleSampleInput)
	assert.DeepEqual(t, []string{"1000", "2000\r\n3000"}, split)
}

func Test_getMaxCalories(t *testing.T) {

	result := getMaxCalories([][]int{{1000, 2000}, {4000, 6000}})
	assert.Equal(t, 10000, result)
}

func Test_splitAll(t *testing.T) {
	result := splitAll(simpleSampleInput)
	assert.DeepEqual(t, [][]int{{1000}, {2000, 3000}}, result)
}

func Test_simple_sample_returns_5000(t *testing.T) {
	result := part1(simpleSampleInput)
	assert.Equal(t, 5000, result)
}

func Test_sample1_returns_24000(t *testing.T) {
	input, _ := puzzleLoader.ReadAllText(filepath.Join("data", "day01_sample1.txt"))
	result := part1(input)
	assert.Equal(t, 24000, result)
}
