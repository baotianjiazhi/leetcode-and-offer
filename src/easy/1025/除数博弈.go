package main

func divisorGame(N int) bool {
	dp := make([]bool, N+1)

	if N == 1 {
		return false
	} else if N == 2 {
		return true
	}
	dp[1] = false
	dp[2] = true

	for i := 3; i <= N; i++ {
		dp[i] = false
		for j := 1; j < i; j++ {
			if (i % j ==0) && (dp[i-j] == false) {
				dp[i] = true
			}
		}
	}

	return dp[N]
}

func main()  {
	print(divisorGame(4))
}
