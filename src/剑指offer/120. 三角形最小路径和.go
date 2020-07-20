package main

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle[0]) == 0 {
		return 0
	}

	for i := len(triangle)-2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min_two(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func min_two(a, b int) int {
	if a > b {
		return b
	}

	return a
}