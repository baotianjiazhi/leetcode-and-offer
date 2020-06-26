package main

var result [][]string

/***
result = []
func backtrack(路径，选择列表) {
	if 满足结束条件 {
		result.add(路径)
	}
	return

	for 选择 in 选择列表 {
		做选择
		backtrack(路径，选择列表)
		撤销选择
	}
}

 */

func backtrack(board [][]bool, path[][]byte) {
	// 结束条件
	if len(path) == len(board) {
		t := make([]string, len(path))
		for k, bs := range path {
			t[k] = string(bs)
		}
		result = append(result, t)
	}

	for i := 0; i < len(board); i++{
		if !isValid(board, len(path), i) {
			continue
		}

		bs := printLine(len(board))
		bs[i] = 'Q'
		board[len(path)][i] = true
		path = append(path, bs)
		backtrack(board, path)
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


// 打印每行默认情况,都是'.'
func printLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}

func solveNQueens(n int) [][]string {
	// 清空变量
	result = [][]string{}
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	backtrack(board, [][]byte{})
	return result
}
