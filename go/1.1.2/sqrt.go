package main

import (
	"fmt"
	"strconv"
)

func absDiff(a, b float64) float64 {
	if a < b {
		return b - a
	}

	return a - b
}

func sqrt(x float64, precision int) float64 {
	if x < 0 {
		return 0
	}

	min, max, cur := 0.0, float64(x), float64(x)/2
	calcPower := cur * cur
	lastPower := -1.0
	pre, _ := strconv.ParseFloat(fmt.Sprintf("1e-%d", precision), 64)

	for absDiff(calcPower, lastPower) >= pre {
		if calcPower > x {
			max = cur
		} else {
			min = cur
		}
		cur = (min + max) / 2
		lastPower = calcPower
		calcPower = cur * cur
	}

	return cur
}

func main() {
	fmt.Println(sqrt(2, 10))
}
