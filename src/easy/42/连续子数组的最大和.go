package main

import "math"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for k := range nums {
		if k == 0 {
			dp[k] = nums[k]
			continue
		}
		dp[k] = int(math.Max(float64(nums[k]), float64(nums[k]+dp[k-1])))

	}

	maxval := dp[0]
	for i :=0; i < len(dp); i++ {
		if maxval < dp[i] {
			maxval = dp[i]
		}
	}

	return maxval
}
