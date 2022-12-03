package puzzleLoader

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
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

// Read the contents of a file as a string
func ReadAllText(path string) (string, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}

// Type of input file
type AoCInput string

const (
	// The real data from the day's puzzle
	Real AoCInput = ""
	// First sample data provided in the puzzle specs
	Sample1 AoCInput = "_sample1"
	// Second sample data provided in the puzzle specs
	Sample2 AoCInput = "_sample2"
	// Third sample data provided in the puzzle specs
	Sample3 AoCInput = "_sample3"
)

func getFileName(year int, day int, input AoCInput) string {
	return filepath.Join(fmt.Sprint(year), fmt.Sprintf("day%02d", day), fmt.Sprintf("day%02d%s.txt", day, input))
}

// Load input puzzle data lines as a string slice
func LoadLines(year int, day int, input AoCInput) ([]string, error) {
	fn := getFileName(year, day, input)
	log.Printf("Loading lines from %s\n", fn)
	return ReadAllLines(fn)
}

// Load input puzzle data as a single data string
func LoadString(year int, day int, input AoCInput) (string, error) {
	fn := getFileName(year, day, input)
	log.Printf("Loading string from %s\n", fn)
	return ReadAllText(fn)
}
