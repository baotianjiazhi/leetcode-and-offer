package main

import (
	"fmt"
	"math"
)

func deleteAndEarn(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(nums[i] * i + dp[i-2]), float64(dp[i-1])))
	}

	return dp[len(nums)-1]
}

func main() {
	a := []int{1, 2, 4, 5, 6 ,8, 1, 3}
	fmt.Println(deleteAndEarn(a))
}