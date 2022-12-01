package day01

import (
	"log"
	"strconv"
	"strings"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
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

func getMaxCalories(input [][]int) int {
	sums := lo.Map(input, func(item []int, index int) int {
		return lo.Sum(item)
	})
	return lo.Max(sums)
}

func getSumOf3BestCalories(input [][]int) int {
	var podium = []int{0, 0, 0}

	// Keeping this one manual as it will probably be faster than doing a sort + take 3
	// equivalent with lo.
	for _, calories := range input {
		sum := lo.Sum(calories)
		if sum >= podium[0] {
			podium = []int{sum, podium[0], podium[1]}
		} else if sum >= podium[1] {
			podium = []int{podium[0], sum, podium[1]}
		} else if sum >= podium[2] {
			podium = []int{podium[0], podium[1], sum}
		}
	}
	return lo.Sum(podium)
}

func part1(input string) int {
	return getMaxCalories(splitAll(input))
}

func part2(input string) int {
	return getSumOf3BestCalories(splitAll(input))
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadString(2022, 1, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Invalid input")
	}

	return part1(input)
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadString(2022, 1, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Invalid input")
	}

	return part2(input)
}
