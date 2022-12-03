package day09

import (
	"fmt"
	"github.com/julienadam/adventofcode2022/puzzleLoader"
	"golang.org/x/exp/maps"
	"log"
	"strconv"
	"strings"
)

func enumeratePermutationsRec[T any](a []T, size int, n int, action func([]T)) {
	// if size becomes 1 then prints the obtained
	// permutation
	if size == 1 {
		action(a)
	}

	for i := 0; i < size; i++ {
		enumeratePermutationsRec(a, size-1, n, action)

		// if size is odd, swap 0th i.e (first) and
		// (size-1)th i.e (last) element
		if size%2 == 1 {
			temp := a[0]
			a[0] = a[size-1]
			a[size-1] = temp
		} else {
			// If size is even, swap ith and
			// (size-1)th i.e (last) element
			temp := a[i]
			a[i] = a[size-1]
			a[size-1] = temp
		}
	}
}

func EnumeratePermutations[T any](a []T, action func([]T)) {
	enumeratePermutationsRec(a, len(a), len(a), action)
}

func pathDistance(locs []int, dists map[int]int) int {
	total := 0
	for i := 1; i < len(locs); i++ {
		prev := locs[i-1]
		curr := locs[i]
		total += dists[prev|curr]
	}
	return total
}

func printLocs(locs []int, dist int) {
	fmt.Printf("New best route found : ")
	for _, i := range locs {
		fmt.Printf("%d ->", i)
	}
	fmt.Printf("(%d)\n", dist)
}

func FindShortestPath(locs []int, segments map[int]int) int {
	min := 0xFFFFFFFF
	EnumeratePermutations(locs, func(perm []int) {
		dist := pathDistance(locs, segments)
		if dist < min {
			printLocs(locs, dist)
			min = dist
		}
	})

	return min
}

func loadData(lines []string) (map[int]int, []int) {
	result := make(map[int]int)
	id := 1
	namedLocs := make(map[string]int)
	for _, line := range lines {
		split := strings.Split(line, " ")
		from := split[0]
		fromId, ok := namedLocs[from]
		if !ok {
			fromId = id
			namedLocs[from] = fromId
			id = id << 1
		}

		to := split[2]
		toId, ok := namedLocs[to]
		if !ok {
			toId = id
			namedLocs[to] = toId
			id = id << 1
		}

		dist, _ := strconv.Atoi(split[4])
		result[fromId|toId] = dist
	}

	return result, maps.Values(namedLocs)
}

func LoadAndSolvePart1() int {
	input, err := puzzleLoader.LoadLines(2015, 9, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	distances, locs := loadData(input)
	return FindShortestPath(locs, distances)
}

func FindLongestPath(locs []int, segments map[int]int) int {
	max := 0
	EnumeratePermutations(locs, func(perm []int) {
		dist := pathDistance(locs, segments)
		if dist > max {
			printLocs(locs, dist)
			max = dist
		}
	})

	return max
}

func LoadAndSolvePart2() int {
	input, err := puzzleLoader.LoadLines(2015, 9, puzzleLoader.Real)
	if err != nil {
		log.Fatal("Could not load data")
	}

	distances, locs := loadData(input)
	return FindLongestPath(locs, distances)
}
