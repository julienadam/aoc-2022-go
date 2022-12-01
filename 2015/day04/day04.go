package day04

import (
	"crypto/md5"
	"fmt"
	"sync"
)

func hashHas5Zeroes(input string) bool {
	bytes := []byte(input)
	hash := md5.Sum(bytes)
	return hash[0] == 0 && hash[1] == 0 && (hash[2]&0xF0) == 0
}

func mineAdventCoin(input string) int {
	i := 0
	for {
		if hashHas5Zeroes(fmt.Sprintf("%s%d", input, i)) {
			return i
		}
		i++
	}
}

func hashHas6Zeroes(input string) bool {
	bytes := []byte(input)
	hash := md5.Sum(bytes)
	return hash[0] == 0 && hash[1] == 0 && hash[2] == 0
}

func threadMiner(input string, start int, step int, wg *sync.WaitGroup, result *int, stop *bool) {
	i := start
	for {
		if *stop {
			break
		}
		if hashHas6Zeroes(fmt.Sprintf("%s%d", input, i)) {
			*result = i
			wg.Done()
			break
		}
		i += step
	}
}

func mineAdventCoinAdvanced(input string) int {
	var wg sync.WaitGroup
	wg.Add(1)
	found := false
	result := 0

	for i := 0; i < 8; i++ {
		go threadMiner(input, i, 8, &wg, &result, &found)
	}

	wg.Wait()
	found = true
	return result
}

func LoadAndSolvePart1() int {
	return mineAdventCoin("iwrupvqb")
}

func LoadAndSolvePart2() int {
	return mineAdventCoinAdvanced("iwrupvqb")
}
