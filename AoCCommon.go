package AoCCommon

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Read all lines into a slice.
// Probably enough for reasonable quantities of data.
// An iterator-like pattern would be preferred for a larger data set.
func ReadAllLines(path string) ([]string, error) {
	fileObject, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fileObject)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
}

type AoCInput string

const (
	Real    AoCInput = ""
	Sample1 AoCInput = "_sample1"
	Sample2 AoCInput = "_sample2"
	Sample3 AoCInput = "_sample3"
)

func LoadData(day int, input AoCInput) ([]string, error) {
	fn := filepath.Join("data", fmt.Sprintf("aoc_2022_day%02d%s.txt", day, input))
	fmt.Printf("Loading data from %s", fn)
	return ReadAllLines(fn)
}
