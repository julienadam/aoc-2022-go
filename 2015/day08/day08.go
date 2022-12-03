package day08

import (
	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
	"log"
)

func CountCharsInMemory(input string) int {
	numChars := 0
	isEscape := false
	skip := 0
	for i, c := range input {
		// Skip first and last chars, as they are quotes
		if i == 0 || i == len(input)-1 {
			continue
		}

		// Skip until skip is 0
		if skip > 0 {
			skip--
			continue
		}

		if isEscape {
			if c == '\\' || c == '"' {
				isEscape = false
				numChars++
			} else if c == 'x' {
				isEscape = false
				skip = 2
				numChars++
			}
		} else if c == '\\' {
			isEscape = true
		} else {
			numChars++
		}
	}

	return numChars
}

func TotalCharsMinusTotalMemory(input []string) int {
	totalChars := lo.SumBy(input, func(line string) int {
		return len(line)
	})
	totalMem := lo.SumBy(input, func(line string) int {
		return CountCharsInMemory(line)
	})
	return totalChars - totalMem
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2015, 8, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return TotalCharsMinusTotalMemory(input)
}

func Escape(input string) string {
	result := ""
	for _, r := range input {
		if r == '"' {
			result += "\\\""
		} else if r == '\\' {
			result += "\\\\"
		} else {
			result += string(r)
		}
	}

	return "\"" + result + "\""
}

func TotalEncodedCharsMinusInitialChars(input []string) int {
	totalChars := lo.SumBy(input, func(line string) int {
		return len(line)
	})
	totalEncoded := lo.SumBy(input, func(line string) int {
		return len(Escape(line))
	})
	return totalEncoded - totalChars
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadLines(2015, 8, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return TotalEncodedCharsMinusInitialChars(input)
}
