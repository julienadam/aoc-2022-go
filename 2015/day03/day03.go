package day03

import (
	"log"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
)

type Point struct {
	x int
	y int
}

func deliver(input string) int {
	visited := make(map[Point]bool)
	visited[Point{0, 0}] = true

	x, y := 0, 0
	for _, r := range input {
		switch r {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		default:
			log.Fatalf("Should not encounter this rune %c", r)
		}

		visited[Point{x, y}] = true
	}

	return len(visited)
}

func deliverWithRoboSanta(input string) int {
	visited := make(map[Point]bool)
	visited[Point{0, 0}] = true

	santaPos, roboPos := Point{0, 0}, Point{0, 0}
	currentDeliveryAgent := &santaPos
	for _, r := range input {
		switch r {
		case '>':
			currentDeliveryAgent.x++
		case '<':
			currentDeliveryAgent.x--
		case '^':
			currentDeliveryAgent.y++
		case 'v':
			currentDeliveryAgent.y--
		default:
			log.Fatalf("Should not encounter this rune %c", r)
		}

		visited[Point{currentDeliveryAgent.x, currentDeliveryAgent.y}] = true

		if currentDeliveryAgent == &santaPos {
			currentDeliveryAgent = &roboPos
		} else {
			currentDeliveryAgent = &santaPos
		}
	}

	return len(visited)
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadString(2015, 03, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Failed to load data")
	}

	return deliver(input)
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadString(2015, 03, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Failed to load data")
	}

	return deliverWithRoboSanta(input)
}
