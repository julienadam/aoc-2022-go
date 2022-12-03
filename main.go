package main

import (
	"fmt"

	"github.com/bradhe/stopwatch"
	"github.com/julienadam/adventofcode2022/2015/day09"
)

func main() {
	sw := stopwatch.Start()
	result := day09.LoadAndSolvePart1()
	sw.Stop()
	fmt.Printf("%d\nDone in %dms", result, sw.Milliseconds())
}
