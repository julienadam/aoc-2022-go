package day02

import (
	"log"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
)

type rps int

const (
	rock     rps = 1
	paper    rps = 2
	scissors rps = 3
)

type round int

const (
	win  round = 6
	draw round = 3
	loss round = 0
)

func charToRps(char rune) rps {
	switch char {
	case 'A':
		return rock
	case 'X':
		return rock
	case 'B':
		return paper
	case 'Y':
		return paper
	case 'C':
		return scissors
	case 'Z':
		return scissors
	default:
		log.Fatalf("Invalid input char %c", char)
	}
	return 0
}

func getPointsForResult(opponent rps, me rps) round {
	if opponent == me {
		return draw
	}
	if opponent == rock {
		if me == paper {
			return win
		} else {
			return loss
		}
	} else if opponent == paper {
		if me == scissors {
			return win
		} else {
			return loss
		}
	} else if opponent == scissors {
		if me == rock {
			return win
		} else {
			return loss
		}
	}

	log.Fatalf("Should not happen")
	return -1
}

func scoreRound(opponent rps, me rps) int {
	return int(getPointsForResult(opponent, me)) + int(me)
}

func parseRound(input string) (rps, rps) {
	return charToRps(rune(input[0])), charToRps(rune(input[2]))
}

func scoreAllRounds(input []string) int {
	return lo.SumBy(input, func(line string) int {
		return scoreRound(parseRound(line))
	})
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2022, 02, puzzleLoader.Real)
	if err != nil {
		log.Fatalf("Could not load data")
	}

	return scoreAllRounds(input)
}