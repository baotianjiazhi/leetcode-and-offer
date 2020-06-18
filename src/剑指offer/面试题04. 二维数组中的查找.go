package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0{
		return false
	}

	if len(matrix[0]) == 0 {
		return false
	}
	columns := len(matrix[0])-1
	row := len(matrix)-1
	column := 0
	for row >= 0 && column <= columns {
		temp := matrix[row][column]
		if temp == target {
			return true
		} else if temp < target {
			column += 1
		} else if temp > target {
			row -= 1
		}
	}

	return false
}
