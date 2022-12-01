package day01

import (
	"log"
	"strconv"
	"strings"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
)

func splitByElf(input string) []string {
	return strings.Split(input, "\r\n\r\n")
}

func splitSingleElfItems(input string) []int {
	items := strings.Split(input, "\r\n")
	var result []int = []int{}

	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		calorie, err := strconv.Atoi(trimmed)
		if err != nil {
			log.Fatalf("Could not parse calorie %s", trimmed)
		}
		result = append(result, calorie)
	}

	return result
}

func splitAll(input string) [][]int {
	var result [][]int
	for _, s := range splitByElf(input) {
		result = append(result, splitSingleElfItems(s))
	}
	return result
}

func Sum(ints []int) int {
	sum := 0
	for _, val := range ints {
		sum += val
	}
	return sum
}

func getMaxCalories(input [][]int) int {
	max := 0
	for _, calories := range input {
		sum := Sum(calories)
		if sum > max {
			max = sum
		}
	}
	return max
}

func part1(input string) int {
	return getMaxCalories(splitAll(input))
}

func LoadAndSolvePart1() int {

	input, err := puzzleLoader.LoadString(2022, 1, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Invalid input")
	}

	return part1(input)
}
