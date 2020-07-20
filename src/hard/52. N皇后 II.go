package main

var res int
func backtrack_2(board [][]bool, path[][] int) {
	// 结束条件
	if len(path) == len(board) {
		res += 1
	}

	for i := 0; i < len(board); i++{
		if !isValid(board, len(path), i) {
			continue
		}
		bs := printLine(len(board))
		bs[i] = 1
		path = append(path, bs)
		board[len(path)][i] = true
		backtrack_2(board, path)
		path = path[:len(path)-1]
		board[len(path)][i] = false
	}
}

func isValid(board [][]bool, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == true {
			return false
		}
	}

	for j := 0; j < col; j ++ {
		if board[row][j] {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}

	for i, j := row, col; i >=0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}

	return true
}

func printLine(n int) []int {
	bs := make([]int, n)
	for i := 0; i < n; i++ {
		bs[i] = 0
	}
	return bs
}

func totalNQueens(n int) int {
	res := 0
	board := make([][]bool, n)
	for i := 0; i < len(board); i++ {
		board[i] = make([]bool, n)
	}

	path := make([][]int, n)
	backtrack_2(board, path)

	return res
}

