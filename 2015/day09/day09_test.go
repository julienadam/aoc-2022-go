package day09

import (
	"github.com/samber/lo"
	"gotest.tools/v3/assert"
	"strconv"
	"testing"
)

func Test_enumerate_permutations_works(t *testing.T) {
	result := ""
	expected := "123\n213\n312\n132\n231\n321\n"
	EnumeratePermutations([]int{1, 2, 3}, func(permutation []int) {
		for _, value := range permutation {
			result += strconv.Itoa(value)
		}
		result += "\n"
	})

	assert.Equal(t, expected, result)
}

func Test_load_data_works_for_sample(t *testing.T) {
	input := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	result, locs := loadData(input)

	assert.Equal(t, 3, len(result))
	assert.Equal(t, 464, result[3])
	assert.Equal(t, 518, result[5])
	assert.Equal(t, 141, result[6])

	assert.Equal(t, 3, len(locs))
	assert.Equal(t, true, lo.Contains(locs, 1))
	assert.Equal(t, true, lo.Contains(locs, 2))
	assert.Equal(t, true, lo.Contains(locs, 4))
}

func Test_sample_shortest_dist_was_605(t *testing.T) {
	input := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	distances, locs := loadData(input)
	assert.Equal(t, 605, FindShortestPath(locs, distances))
}

func Test_sample_max_dist_was_982(t *testing.T) {
	input := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	distances, locs := loadData(input)
	assert.Equal(t, 982, FindLongestPath(locs, distances))
}
