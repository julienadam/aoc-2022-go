package day06

import (
	"log"
	"regexp"
	"strconv"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
)

func turnOnThrough(grid *[1000][1000]bool, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] = true
		}
	}
}

func turnOffThrough(grid *[1000][1000]bool, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] = false
		}
	}
}

func toggleThrough(grid *[1000][1000]bool, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] = !grid[x][y]
		}
	}
}

type SwitchAction int

const (
	on     SwitchAction = iota
	off    SwitchAction = iota
	toggle SwitchAction = iota
)

type Instruction struct {
	X1     int
	Y1     int
	X2     int
	Y2     int
	Action SwitchAction
}

func strToSwitchAction(input string) SwitchAction {
	switch input {
	case "turn on":
		return on
	case "turn off":
		return off
	case "toggle":
		return toggle
	}

	log.Fatalf("Invalid switch action %s", input)

	return -1
}

func parseInstruction(line string) Instruction {
	re := regexp.MustCompile("(turn on|turn off|toggle) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)")
	matches := re.FindStringSubmatch(line)
	x1, _ := strconv.Atoi(matches[2])
	y1, _ := strconv.Atoi(matches[3])
	x2, _ := strconv.Atoi(matches[4])
	y2, _ := strconv.Atoi(matches[5])

	return Instruction{x1, y1, x2, y2, strToSwitchAction(matches[1])}
}

func applyInstructions(grid *[1000][1000]bool, instructions []Instruction) {
	for _, instr := range instructions {
		switch instr.Action {
		case on:
			turnOnThrough(grid, instr.X1, instr.Y1, instr.X2, instr.Y2)
		case off:
			turnOffThrough(grid, instr.X1, instr.Y1, instr.X2, instr.Y2)
		case toggle:
			toggleThrough(grid, instr.X1, instr.Y1, instr.X2, instr.Y2)
		}
	}
}

func countLightsOn(grid [1000][1000]bool) int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				count++
			}
		}
	}

	return count
}

func loadInstructions() []Instruction {
	input, err := puzzleLoader.LoadLines(2015, 06, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return lo.Map(input, func(line string, index int) Instruction {
		return parseInstruction(line)
	})
}

func LoadAndSolvePart1() int {
	instructions := loadInstructions()
	grid := [1000][1000]bool{}
	applyInstructions(&grid, instructions)
	return countLightsOn(grid)
}

func increaseThrough(grid *[1000][1000]int, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y]++
		}
	}
}

func decreaseThrough(grid *[1000][1000]int, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if grid[x][y] > 0 {
				grid[x][y]--
			}
		}
	}
}

func increaseMoreThrough(grid *[1000][1000]int, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] += 2
		}
	}
}

func applyUpdatedInstructions(grid *[1000][1000]int, instructions []Instruction) {
	for _, instr := range instructions {
		switch instr.Action {
		case on:
			increaseThrough(grid, instr.X1, instr.Y1, instr.X2, instr.Y2)
		case off:
			decreaseThrough(grid, instr.X1, instr.Y1, instr.X2, instr.Y2)
		case toggle:
			increaseMoreThrough(grid, instr.X1, instr.Y1, instr.X2, instr.Y2)
		}
	}
}

func countBrightness(grid [1000][1000]int) int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			count += grid[x][y]
		}
	}

	return count
}

func LoadAndSolvePart2() int {
	instructions := loadInstructions()
	grid := [1000][1000]int{}

	applyUpdatedInstructions(&grid, instructions)
	return countBrightness(grid)
}
