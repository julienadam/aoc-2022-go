package day04

import (
	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/julienadam/adventofcode2022/util"
	"github.com/samber/lo"
	"log"
	"strconv"
	"strings"
)

func parseRanges(s string) (util.Range[int], util.Range[int]) {
	pairs := strings.Split(s, ",")
	pairA := strings.Split(pairs[0], "-")
	pairB := strings.Split(pairs[1], "-")
	a, _ := strconv.Atoi(pairA[0])
	b, _ := strconv.Atoi(pairA[1])
	c, _ := strconv.Atoi(pairB[0])
	d, _ := strconv.Atoi(pairB[1])
	return util.Range[int]{a, b}, util.Range[int]{c, d}
}

func findCompletelyOverlappingPairs(input []string) int {
	return lo.CountBy(input, func(line string) bool {
		r1, r2 := parseRanges(line)
		return r1.ContainsOrIsContainedIn(r2)
	})
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2022, 4, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return findCompletelyOverlappingPairs(input)
}

func findPartiallyOverlappingPairs(input []string) int {
	return lo.CountBy(input, func(line string) bool {
		r1, r2 := parseRanges(line)
		return r1.Overlaps(r2)
	})
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadLines(2022, 4, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return findPartiallyOverlappingPairs(input)
}
