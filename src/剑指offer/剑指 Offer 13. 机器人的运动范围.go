package main

func movingCount(m int, n int, k int) int {
	if m == 0 && n == 0 {
		return 0
	}

	arr := make([][]bool, m)
	for i, _ := range arr {
		arr[i] = make([]bool, n)
	}

	for i := 0; i < len(arr); i++{
		for j := 0; j <len(arr[0]); j++ {
			arr[i][j] = false
		}
	}

	//arr[0][0] = true
	count := Countersum(m, n, 0, 0, k, arr)

	return count
}

var sum int
func Countersum(m, n, i, j, k int, arr [][]bool) int {
	sum = 0
	if check(m, n, i, j, k, arr) {
		arr[i][j] = true
		sum = 1 + Countersum(m, n, i + 1, j, k, arr) + Countersum(m, n, i - 1, j, k, arr) + Countersum(m, n, i, j+1, k, arr) + Countersum(m, n, i, j-1, k, arr)
	}

	return sum
}

func check(m, n, i, j, k int, arr [][]bool) bool {
	if (i >= 0 && i < m && j >= 0 && j < n && getDigitSum(i) + getDigitSum(j) <= k && !arr[i][j]) {
		return true
	}

	return false
}

func getDigitSum(n int) (sum int) {

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}