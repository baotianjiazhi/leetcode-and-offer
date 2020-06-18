package main

func largest1BorderedSquare(grid [][]int) int {
	gx := make([][]int, len(grid)+1)
	gy := make([][]int, len(grid)+1)

	for i := 0; i <= len(grid); i++ {
		gx[i] = make([]int, len(grid[0])+1)
		gy[i] = make([]int, len(grid[0])+1)
	}


}
