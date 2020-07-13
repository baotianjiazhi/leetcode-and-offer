package main

func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	a, b, c := 0, 0, 0
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a] * 2, dp[b] * 3, dp[c] * 5
		dp[i] = min(n2, n3, n5)
		if dp[i] == n2 {
			a++
		}

		if dp[i] == n3 {
			b++
		}

		if dp[i] == n5 {
			c++
		}
	}

	return dp[n-1]
}

func min(a, b, c int) (int) {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}
