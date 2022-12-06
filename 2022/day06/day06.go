package day06

import (
	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
	"log"
)

func findStartOfPacketMarker(input string) int {
	a := input[0]
	b := input[1]
	c := input[2]

	for i := 3; i < len(input); i++ {
		r := input[i]
		if r != a && r != b && r != c && a != b && a != c && b != c {
			return i + 1
		} else {
			a = input[i-2]
			b = input[i-1]
			c = input[i]
		}
	}

	log.Fatalf("Could not find marker")
	return 0
}

func findStartOfMessageMarker(input string) int {

	for i := 13; i < len(input); i++ {
		r := []byte(input[i-13 : i+1])
		uniques := lo.FindUniques(r)
		if len(r) == len(uniques) {
			return i + 1
		}
	}

	log.Fatalf("Could not find marker")
	return 0
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadString(2022, 6, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return findStartOfPacketMarker(input)
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadString(2022, 6, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return findStartOfMessageMarker(input)
}
