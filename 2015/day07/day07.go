package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/julienadam/adventofcode2022/puzzleLoader"
)

func SplitLeftRight(instruction string) (string, string) {
	s := strings.Split(instruction, " -> ")
	return strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
}

func atouint16(input string) uint16 {
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Could not parse %s into a uint16", input)
	}
	return uint16(i)
}

func getValOrFindSignal(input string, layout map[string]string, memo map[string]uint16) uint16 {
	i, err := strconv.Atoi(input)
	if err != nil {
		return findSignalForWire(input, layout, memo)
	} else {
		return uint16(i)
	}
}

func findSignalForWire(wire string, layout map[string]string, memo map[string]uint16) uint16 {
	memoed, ok := memo[wire]
	if ok {
		return memoed
	}

	connection := layout[wire]
	connectionParts := strings.Split(connection, " ")
	result := uint16(0)

	switch len(connectionParts) {
	case 1:
		result = getValOrFindSignal(connectionParts[0], layout, memo)
	case 2: // This is a NOT xyz
		l := connectionParts[1]
		result = ^(findSignalForWire(l, layout, memo))
	case 3:
		op := connectionParts[1]
		l := getValOrFindSignal(connectionParts[0], layout, memo)
		switch op {
		case "LSHIFT":
			r := connectionParts[2]
			result = l << atouint16(r)
		case "RSHIFT":
			r := connectionParts[2]
			result = l >> atouint16(r)
		case "AND":
			r := getValOrFindSignal(connectionParts[2], layout, memo)
			result = l & r
		case "OR":
			r := getValOrFindSignal(connectionParts[2], layout, memo)
			result = l | r
		default:
			log.Fatalf("Unknown operator %s", connectionParts[1])
		}
	default:
		log.Fatalf("Unparsable instruction %s, more parts than expected", connection)
	}

	memo[wire] = result
	return result
}

func loadInstructions() map[string]string {
	input, err := puzzleLoader.LoadLines(2015, 07, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	layout := make(map[string]string)
	for _, line := range input {
		l, r := SplitLeftRight(line)
		layout[r] = l
	}

	return layout
}

func LoadAndSolvePart1() int {
	layout := loadInstructions()
	return int(findSignalForWire("a", layout, make(map[string]uint16)))
}

func LoadAndSolvePart2() int {
	layout := loadInstructions()
	// run once to get the value for "a"
	a := findSignalForWire("a", layout, make(map[string]uint16))
	// override "b" with value from "a" found just above
	layout["b"] = fmt.Sprintf("%d", a)
	// run again with the modified input
	return int(findSignalForWire("a", layout, make(map[string]uint16)))
}
