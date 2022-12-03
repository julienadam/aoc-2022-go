package main

import (
	"fmt"

	"github.com/bradhe/stopwatch"
	"github.com/julienadam/adventofcode2022/2015/day08"
)

func main() {
	sw := stopwatch.Start()
	result := day08.LoadAndSolvePart2()
	sw.Stop()
	fmt.Printf("%d\nDone in %dms", result, sw.Milliseconds())
}
