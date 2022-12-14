package day03

import (
	"log"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
)

func SplitRucksack(rucksack string) (string, string) {
	half := len(rucksack) / 2
	return rucksack[0:half], rucksack[half:]
}

func FindCommonItemInCompartments(a string, b string) rune {
	itemsInA := make(map[rune]bool)
	for _, item := range a {
		itemsInA[item] = true
	}

	for _, item := range b {
		_, ok := itemsInA[item]
		if ok {
			return item
		}
	}

	log.Fatalf("Could not find common item in %s and %s", a, b)
	return rune(0)
}

func getPriority(item rune) int {
	if int(item) >= 'a' {
		return int(item) - int('a') + 1
	} else {
		return int(item) - int('A') + 27
	}
}

func SumOfPriorities(rucksacks []string) int {
	return lo.SumBy(rucksacks, func(rucksack string) int {
		a, b := SplitRucksack(rucksack)
		common := FindCommonItemInCompartments(a, b)
		return getPriority(common)
	})
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2022, 03, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Could not load data")
	}

	return SumOfPriorities(input)
}

func GroupRuckSacks(rucksacks []string) [][]string {
	return lo.Chunk(rucksacks, 3)
}

func FindBadge(rucksacks []string) rune {
	itemsA := make(map[rune]bool)
	commonAB := make(map[rune]bool)
	a := rucksacks[0]
	b := rucksacks[1]
	c := rucksacks[2]

	for _, item := range a {
		itemsA[item] = true
	}

	for _, item := range b {
		_, ok := itemsA[item]
		if ok {
			commonAB[item] = true
		}
	}

	for _, item := range c {
		_, ok := commonAB[item]
		if ok {
			return item
		}
	}

	log.Fatalf("Should not happen")
	return rune(0)
}

func SumOfBadgePriorities(rucksacks []string) int {
	groups := GroupRuckSacks(rucksacks)
	return lo.SumBy(groups, func(group []string) int {
		common := FindBadge(group)
		return getPriority(common)
	})
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadLines(2022, 03, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Could not load data")
	}

	return SumOfBadgePriorities(input)
}
