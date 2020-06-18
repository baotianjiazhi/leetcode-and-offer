package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	f := make([]int, len(cost))
	f[0] = cost[0]
	f[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		f[i] = cost[i] + int(math.Min(float64(f[i-1]), float64(f[i-2])))
	}

	minv := 0
	if f[len(f)-2] < f[len(f)-1] {
		minv = f[len(f) -2]
	} else {
		minv = f[len(f) - 1]
	}
	return  minv
}

func main() {
	a := []int{0, 0, 0, 1}

	fmt.Println(minCostClimbingStairs(a))
}