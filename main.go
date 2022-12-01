package main

import (
	"fmt"

	"github.com/bradhe/stopwatch"
	"github.com/julienadam/adventofcode2022/2015/day04"
)

func main() {
	sw := stopwatch.Start()
	day04.LoadAndSolvePart2()
	sw.Stop()
	fmt.Printf("Done in %dms", sw.Milliseconds())
}
