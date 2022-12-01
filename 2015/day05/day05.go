package day05

import (
	"log"
	"strings"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
)

func isVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
		return true
	}
	return false
}

func hasTwoConsecutiveEqualLetters(input string) bool {
	prev := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] == prev {
			return true
		} else {
			prev = input[i]
		}
	}
	return false
}

func isNice(input string) bool {
	if strings.Contains(input, "xy") {
		return false
	}
	if strings.Contains(input, "ab") {
		return false
	}
	if strings.Contains(input, "cd") {
		return false
	}
	if strings.Contains(input, "pq") {
		return false
	}

	numVowels := lo.CountBy([]rune(input), isVowel)

	conseqRunes := hasTwoConsecutiveEqualLetters(input)

	return numVowels >= 3 && conseqRunes
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2015, 05, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return lo.CountBy(input, isNice)
}

func hasPairWithSingleLetterBetween(input string) bool {
	if len(input) < 3 {
		return false
	}

	prevprev := input[0]

	for i := 2; i < len(input); i++ {
		if input[i] == prevprev {
			return true
		}
		prevprev = input[i-1]
	}

	return false
}

func hasRepeatedNonOverlappingPair(input string) bool {
	if len(input) < 4 {
		return false
	}

	for i := 0; i < len(input)-3; i++ {
		pair := input[i : i+2]
		rest := input[i+2:]

		if strings.Contains(rest, pair) {
			return true
		}
	}

	return false
}

func isExtraNice(input string) bool {
	return hasPairWithSingleLetterBetween(input) && hasRepeatedNonOverlappingPair(input)
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadLines(2015, 05, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return lo.CountBy(input, isExtraNice)
}
