package day04

import (
	"crypto/md5"
	"fmt"
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

func mineAdventCoinAdvanced(input string) int {
	i := 0
	for {
		if hashHas6Zeroes(fmt.Sprintf("%s%d", input, i)) {
			return i
		}
		i++
	}
}

func LoadAndSolvePart1() int {
	return mineAdventCoin("iwrupvqb")
}

func LoadAndSolvePart2() int {
	return mineAdventCoinAdvanced("iwrupvqb")
}
