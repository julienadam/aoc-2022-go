package day07

import (
	"fmt"
	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"log"
	"strconv"
	"strings"
)

func sumOfSmallSizeDirectories(input []string) int {
	dirSizes := getDirSizes(input)
	printSizes(dirSizes)

	sizes := maps.Values(dirSizes)
	smallOnes := getFoldersUnder100000(sizes)
	return lo.Sum(smallOnes)
}

func getFoldersUnder100000(sizes []int) []int {
	smallOnes := lo.Filter(sizes, func(size, _ int) bool {
		return size <= 100000
	})
	return smallOnes
}

func printSizes(input map[string]int) {
	for k, v := range input {
		fmt.Printf("%s %d\n", k, v)
	}
}

func getDirSizes(input []string) map[string]int {
	sizes := make(map[string]int)

	pwdParts := []string{}
	pwd := "/"
	for i, line := range input {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(line, "$ cd ") {
			cd := line[5:]
			pwdParts = updateCurrentPath(cd, pwdParts)
			pwd = buildPath(pwdParts)
			ensureDir(sizes, pwd)
		} else if line == "$ ls" {
			// nothing to do ?
		} else {
			if strings.HasPrefix(line, "dir ") {
				// nothing to do ?
			} else {
				cwd := "/"
				addFileSizeToDir(line, sizes, cwd)
				for i, _ := range pwdParts {
					cwd = buildPath(pwdParts[0 : i+1])
					addFileSizeToDir(line, sizes, cwd)
				}
			}
		}
	}

	return sizes
}

func addFileSizeToDir(line string, sizes map[string]int, pwd string) {
	split := strings.Split(line, " ")
	size, _ := strconv.Atoi(split[0])
	sizes[pwd] = sizes[pwd] + size
}

func buildPath(pwdParts []string) string {
	return "/" + strings.Join(pwdParts, "/")
}

func updateCurrentPath(cd string, pwdParts []string) []string {
	if cd == ".." {
		pwdParts = pwdParts[0 : len(pwdParts)-1]
	} else {
		pwdParts = append(pwdParts, cd)
	}
	return pwdParts
}

func ensureDir(sizes map[string]int, pwd string) {
	_, ok := sizes[pwd]
	if !ok {
		sizes[pwd] = 0
	}
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2022, 7, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return sumOfSmallSizeDirectories(input)
}

func smallestDirToDelete(input []string) int {
	sizes := getDirSizes(input)
	printSizes(sizes)
	const target = 30_000_000
	const total = 70_000_000
	usedSpace := sizes["/"]
	freeSpace := total - usedSpace
	spaceToFree := target - freeSpace

	candidates := lo.Filter(maps.Keys(sizes), func(key string, _ int) bool {
		v := sizes[key]
		if v >= spaceToFree {
			fmt.Printf("Candidate : %s %d\n", key, v)
			return true
		}
		return false
	})

	candidateSizes := lo.Map(candidates, func(key string, _ int) int {
		return sizes[key]
	})

	return lo.Min(candidateSizes)
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadLines(2022, 7, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	return smallestDirToDelete(input)
}
