package day07

import (
	"gotest.tools/v3/assert"
	"testing"
)

var sample1 = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func Test_sum_of_small_dirs_for_part1_is_95437(t *testing.T) {
	result := sumOfSmallSizeDirectories(sample1)

	assert.Equal(t, 95437, result)
}

func Test_smallest_dir_to_delete_for_sample(t *testing.T) {
	result := smallestDirToDelete(sample1)

	assert.Equal(t, 24933642, result)
}
