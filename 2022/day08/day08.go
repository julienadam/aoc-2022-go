package day08

import (
	"log"
	"strconv"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
)

func isOnTheSides(x int, y int, forest [][]int) bool {
	if x == 0 || y == 0 {
		return true
	}

	if x == len(forest)-1 || y == len(forest[0])-1 {
		return true
	}
	return false
}

func isHigherThanAllTreesUp(forest [][]int, row int, col int) bool {
	val := forest[row][col]
	for cr := row - 1; cr >= 0; cr-- {
		if forest[cr][col] >= val {
			return false
		}
	}
	return true
}

func isHigherThanAllTreesDown(forest [][]int, row int, col int) bool {
	val := forest[row][col]
	for cr := row + 1; cr < len(forest); cr++ {
		if forest[cr][col] >= val {
			return false
		}
	}
	return true
}

func isHigherThanAllTreesLeft(forest [][]int, row int, col int) bool {
	val := forest[row][col]
	for cc := col - 1; cc >= 0; cc-- {
		if forest[row][cc] >= val {
			return false
		}
	}
	return true
}

func isHigherThanAllTreesRight(forest [][]int, row int, col int) bool {
	val := forest[row][col]
	for cc := col + 1; cc < len(forest); cc++ {
		if forest[row][cc] >= val {
			return false
		}
	}
	return true
}

func isHigherThanTreesUpDownLeftOrRight(forest [][]int, row int, col int) bool {
	if isHigherThanAllTreesUp(forest, row, col) {
		return true
	}

	if isHigherThanAllTreesDown(forest, row, col) {
		return true
	}

	if isHigherThanAllTreesLeft(forest, row, col) {
		return true
	}

	if isHigherThanAllTreesRight(forest, row, col) {
		return true
	}

	return false
}

func isVisible(forest [][]int, x int, y int) bool {
	if isOnTheSides(x, y, forest) {
		return true
	}

	if isHigherThanTreesUpDownLeftOrRight(forest, x, y) {
		return true
	}

	return false
}

func findVisibleTrees(forest [][]int) int {

	count := 0
	for r, row := range forest {
		for c, _ := range row {
			if isVisible(forest, r, c) {
				count++
			}
		}
	}
	return count
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2022, 8, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}
	grid := lo.Map(input, func(line string, index int) []int {
		ints := lo.Map([]rune(line), func(char rune, index int) int {
			r, _ := strconv.Atoi(string(char))
			return r
		})
		return ints
	})

	return findVisibleTrees(grid)
}
