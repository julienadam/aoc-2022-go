package day04

import (
	"crypto/md5"
	"fmt"
	"sync"
	"sync/atomic"
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

func threadMiner(input string, wg *sync.WaitGroup, current *int64, stop *bool) {
	defer wg.Done()
	for {
		if *stop {
			break
		}
		i := atomic.AddInt64(current, 1)
		if hashHas6Zeroes(fmt.Sprintf("%s%d", input, i)) {
			fmt.Print(i)
			*stop = true
			break
		}
	}
}

func mineAdventCoinAdvanced(input string) {
	var wg sync.WaitGroup
	const THREADS = 8
	wg.Add(THREADS)
	found := false
	current := int64(0)

	for i := 0; i < THREADS; i++ {
		go threadMiner(input, &wg, &current, &found)
	}

	wg.Wait()
	found = true
}

func LoadAndSolvePart1() int {
	return mineAdventCoin("iwrupvqb")
}

func LoadAndSolvePart2() {
	mineAdventCoinAdvanced("iwrupvqb")
}
