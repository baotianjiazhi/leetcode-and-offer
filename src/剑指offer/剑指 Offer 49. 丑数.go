package main

func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	i := 0
	for len(dp) < n {
		dp = append(dp, dp[i]*2)
		dp = append(dp, dp[i]*3)
		dp = append(dp, dp[i])
	}
}
