package main


var res []int
func spiralOrder(matrix [][]int) []int {
	res = []int{}
	start := 0
	if len(matrix) == 0 {
		return []int{}
	}
	rows := len(matrix)
	columns := len(matrix[0])
	for start * 2 < rows && start * 2 < columns {
		append_data(matrix, start, rows, columns)
		start++
	}

	return res
}


func append_data(matrix [][]int, start, rows, columns int) {

	endY := rows - start -1
	endX := columns - start - 1

	for i := start; i <= endX; i++ {
		res = append(res, matrix[start][i])
	}

	if start < endY {
		for i := start+1; i <= endY; i++ {
			res = append(res, matrix[i][endX])
		}
	}

	if start < endY && start < endX {
		for i := endX-1; i >= start; i-- {
			res = append(res, matrix[endY][i])
		}
	}

	if start < endY - 1 && start < endX{
		for i := endY-1; i > start; i-- {
			res = append(res, matrix[i][start])
		}
	}


}