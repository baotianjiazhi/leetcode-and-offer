package main

type Pair struct {
	First  int
	Second int
}

func stoneGame(piles []int) bool {
	dp := make([][]Pair, len(piles))
	for i := 0; i < len(piles); i++ {
		dp[i] = make([]Pair, len(piles))
	}

	for i := 0; i < len(piles); i++ {
		dp[i][i].First = piles[i]
		dp[i][i].Second = 0
	}

	for l := 2; l <= len(piles); l++ {
		for i := 0; i <= len(piles) - l; i++ {
			j := l + i -1
			left := piles[i] + dp[i+1][j].Second
			right := piles[j] + dp[i][j-1].Second

			if left > right {
				dp[i][j].First = left
				dp[i][j].Second = dp[i+1][j].First
			} else {
				dp[i][j].First = right
				dp[i][j].Second = dp[i][j-1].First
			}
		}
	}

	if (dp[0][len(piles)-1].First - dp[0][len(piles)-1].Second) > 0{
		return true
	} else {
		return false
	}
}
