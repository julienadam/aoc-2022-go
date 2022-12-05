package day05

import (
	stack "github.com/golang-ds/stack/slicestack"
	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
	"log"
	"strconv"
	"strings"
)

func readStackState(lines []string) []stack.SliceStack[rune] {
	// Number of columns
	columns := (len(lines[0]) + 1) / 4
	stacks := make([]stack.SliceStack[rune], columns)

	for i := len(lines) - 2; i >= 0; i-- {
		for pos := 0; pos < columns; pos++ {
			read := rune(lines[i][pos*4+1])
			switch read {
			case ' ':
				break
			default:
				stacks[pos].Push(read)
			}
		}
	}

	return stacks
}

type Move struct {
	Quantity int
	From     int
	To       int
}

func readMoves(lines []string) []Move {
	return lo.Map(lines, func(line string, index int) Move {
		split := strings.Split(line, " ")
		q, _ := strconv.Atoi(split[1])
		f, _ := strconv.Atoi(split[3])
		t, _ := strconv.Atoi(split[5])
		return Move{q, f, t}
	})
}

func applyMove(state []stack.SliceStack[rune], move Move) {
	for i := 0; i < move.Quantity; i++ {
		popped, ok := state[move.From-1].Pop()
		if !ok {
			log.Fatalf("Could not pop from %d", i)
		}
		state[move.To-1].Push(popped)
	}
}

func applyMovesWithMoverAndGetResult(state []stack.SliceStack[rune], moves []Move, mover func([]stack.SliceStack[rune], Move)) string {
	for _, move := range moves {
		mover(state, move)
	}

	result := []rune{}
	for _, st := range state {
		top, _ := st.Top()
		result = append(result, top)
	}

	return string(result)
}

func applyMovesAndGetResultWithCrateMover9000(state []stack.SliceStack[rune], moves []Move) string {
	return applyMovesWithMoverAndGetResult(state, moves, applyMove)
}

func loadStateAndMoves() ([]stack.SliceStack[rune], []Move) {
	input, err := puzzleLoader.LoadString(2022, 5, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	split := strings.Split(input, "\r\n\r\n")
	stackInput := strings.Split(split[0], "\r\n")
	movesInput := strings.Split(split[1], "\r\n")
	state := readStackState(stackInput)
	moves := readMoves(movesInput)

	return state, moves
}

func LoadAndSolvePart1() string {
	state, moves := loadStateAndMoves()
	return applyMovesAndGetResultWithCrateMover9000(state, moves)
}

func applyMoveCrateMover9001(state []stack.SliceStack[rune], move Move) {
	var popped stack.SliceStack[rune]
	for i := 0; i < move.Quantity; i++ {
		p, ok := state[move.From-1].Pop()
		if !ok {
			log.Fatalf("Could not pop from %d", i)
		}
		popped.Push(p)
	}

	for i := 0; i < move.Quantity; i++ {
		p, _ := popped.Pop()
		state[move.To-1].Push(p)
	}
}

func applyMovesAndGetResultWithCrateMover9001(state []stack.SliceStack[rune], moves []Move) string {
	return applyMovesWithMoverAndGetResult(state, moves, applyMoveCrateMover9001)
}

func LoadAndSolvePart2() string {
	state, moves := loadStateAndMoves()
	return applyMovesAndGetResultWithCrateMover9001(state, moves)
}
