package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	if m < 1 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n < 1 {
		return 0
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}


	dp := make([][]int, m+2)
	for i := 0; i < m+2; i++{
		dp[i] = make([]int, n+2)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1{
				dp[i][j] = 1
				continue
			}

			fmt.Println(dp)

			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
