package main


func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	if len(grid[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(grid)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}

	dp[1][1] = grid[0][0]
	for i := 1; i <= len(grid); i++ {
		for j := 1; j <= len(grid[0]); j++ {
			dp[i][j] = max(dp[i-1][j]+grid[i-1][j-1], dp[i][j-1]+grid[i-1][j-1])
		}
	}

	return dp[len(grid)][len(grid[0])]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}