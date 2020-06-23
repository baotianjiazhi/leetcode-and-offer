package main

import "fmt"

func cuttingRope_1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2  {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n;i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			if dp[i-j] * dp[j] > max {
				max = dp[i-j] * dp[j]
			}
		}
		dp[i] = max % 1000000007
	}
	return dp[n]
}

func main() {
	fmt.Println(cuttingRope_1(127))
}