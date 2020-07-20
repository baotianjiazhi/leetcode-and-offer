package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map [int]int{}

	for _, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{v, target-v}
		}

		m[v]++
	}

	return nil
}

func main() {
	m := []int{2, 7, 11, 15}
	s := 9
	fmt.Println(twoSum(m, s))
}
