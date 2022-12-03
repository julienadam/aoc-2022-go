package main

import (
	"fmt"

	"github.com/bradhe/stopwatch"
	"github.com/julienadam/adventofcode2022/2022/day03"
)

func main() {
	sw := stopwatch.Start()
	result := day03.LoadAndSolvePart1()
	sw.Stop()
	fmt.Printf("%d\nDone in %dms", result, sw.Milliseconds())
}
