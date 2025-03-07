package main

import (
	"flag"
	"fmt"
)

func main() {
	minPtr := flag.Float64("min", 0, "the min price")
	maxPtr := flag.Float64("max", 0, "the max price")
	min, max := checkMinMax(*minPtr, *maxPtr)

	delta := max - min
	step := delta / 3
	buyPrc, sellPrc := min+step, max-step
	fmt.Printf("buy/sell => %.02f/%.02f", buyPrc, sellPrc)
}

func checkMinMax(min, max float64) (float64, float64) {
	if min <= 0 {
		panic("illegal min price")
	}
	if max <= 0 {
		panic("illegal max price")
	}
	if min > max {
		return max, min
	}
	return min, max
}
