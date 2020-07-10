package main

import "fmt"

func divingBoard(shorter int, longer int, k int) []int {

	if k == 0 {
		return []int{}
	}
	res := []int{}
	m := map[int]int{}
	for i := 0; i <= k; i ++ {
		if m[i * longer + (k-i) * shorter] < 1 {
			res = append(res, i * longer + (k-i) * shorter)
		}

		if _, ok := m[i * longer + (k-i) * shorter]; !ok {
			m[i * longer + (k-i) * shorter]++
		}
	}

	return res
}

func main() {

	fmt.Println(divingBoard(1, 1, 100))
}