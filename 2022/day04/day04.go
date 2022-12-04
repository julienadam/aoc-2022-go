package day04

import (
	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
	"log"
	"strconv"
	"strings"
)

func isRangeCompletelyInsideOtherRange(a int, b int, c int, d int) bool {
	return (c >= a && d <= b) || (c <= a && d >= b)
}

func parseRange(s string) (int, int, int, int) {
	pairs := strings.Split(s, ",")
	pairA := strings.Split(pairs[0], "-")
	pairB := strings.Split(pairs[1], "-")

	a, _ := strconv.Atoi(pairA[0])
	b, _ := strconv.Atoi(pairA[1])
	c, _ := strconv.Atoi(pairB[0])
	d, _ := strconv.Atoi(pairB[1])
	return a, b, c, d
}

func findCompletelyOverlappingPairs(input []string) int {
	return lo.CountBy(input, func(line string) bool {
		return isRangeCompletelyInsideOtherRange(parseRange(line))
	})
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2022, 4, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return findCompletelyOverlappingPairs(input)
}

func isRangeOverlappingRange(a int, b int, c int, d int) bool {
	return a <= d && b >= c
}

func findPartiallyOverlappingPairs(input []string) int {
	return lo.CountBy(input, func(line string) bool {
		return isRangeOverlappingRange(parseRange(line))
	})
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadLines(2022, 4, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return findPartiallyOverlappingPairs(input)
}
