package day02

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
)

func parsePresentSize(sizes string) (int, int, int) {
	split := strings.Split(sizes, "x")
	l, err1 := strconv.Atoi(split[0])
	w, err2 := strconv.Atoi(split[1])
	h, err3 := strconv.Atoi(split[2])

	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatalf("Could not parse %s", sizes)
	}

	return l, w, h
}

func calculatePaperForPackage(l int, w int, h int) int {
	side1 := l * w
	minSide := side1
	side2 := w * h
	if side2 < minSide {
		minSide = side2
	}
	side3 := h * l
	if side3 < minSide {
		minSide = side3
	}
	return 2*side1 + 2*side2 + 2*side3 + minSide
}

func calculateTotalPaper(input []string) int {
	var paperTotal int = 0
	for _, size := range input {
		l, w, h := parsePresentSize(size)
		paperTotal += calculatePaperForPackage(l, w, h)
	}

	return paperTotal
}

func calculateRibbonForPackage(l int, w int, h int) int {
	in := []int{l, w, h}
	sort.Ints(in)
	a := in[0]
	b := in[1]
	c := in[2]

	wrap := a*2 + b*2
	vol := a * b * c
	return wrap + vol
}

func calculateTotalRibbon(input []string) int {
	var ribbonTotal int = 0
	for _, size := range input {
		l, w, h := parsePresentSize(size)
		ribbonTotal += calculateRibbonForPackage(l, w, h)
	}

	return ribbonTotal
}

func loadPuzzleData() []string {
	lines, err := puzzleLoader.LoadLines(2015, 02, puzzleLoader.Real)
	if err != nil {
		fmt.Println("Error loading data")
		return nil
	} else {
		return lines
	}
}

func SolvePart1() {
	lines := loadPuzzleData()
	result := calculateTotalPaper(lines)
	fmt.Printf("%d\n", result)
}

func SolvePart2() {
	lines := loadPuzzleData()
	result := calculateTotalRibbon(lines)
	fmt.Printf("%d\n", result)
}
