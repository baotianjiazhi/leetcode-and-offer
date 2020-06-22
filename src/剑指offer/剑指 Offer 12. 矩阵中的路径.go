package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 && len(word) != 0 {
		return false
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j ++ {
			if haspath(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}


func haspath(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || j <0 || i == len(board) || j == len(board[0]) {
		return false
	}

	if board[i][j] != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	temp := board[i][j]
	board[i][j] = ' '

	if haspath(board, word, i, j+1, k+1) || haspath(board, word, i+1, j, k+1) || haspath(board, word, i-1, j, k+1) ||
		haspath(board, word, i, j-1, k+1){
		return true
	}
	board[i][j] = temp

	return false
}