package day01

import (
	"fmt"
	"log"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
)

func findFinalFloor(input string) int {
	floor := 0

	for _, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	return floor
}

func findBasementInstructionIndex(input string) int {
	floor := 0

	for i, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
			if floor == -1 {
				return i + 1
			}
		}
	}

	return -1
}

func SolvePart1() {
	input, err := puzzleLoader.LoadString(2015, 1, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Could not load data %s", err)
	} else {
		result := findFinalFloor(input)
		fmt.Printf("%d\n", result)
	}
}

func SolvePart2() {
	input, err := puzzleLoader.LoadString(2015, 1, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Could not load data %s", err)
	} else {
		result := findBasementInstructionIndex(input)
		fmt.Printf("%d\n", result)
	}
}
