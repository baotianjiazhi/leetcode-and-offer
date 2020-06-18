package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	m := map[int]int{}
	for _, v := range nums{
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}

	for k, v := range m {
		if v > 1 {
			return k
		}
	}

	return 0
}

func findRepeatNumber_2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}

			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return 0
}

func main() {
	a := []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(findRepeatNumber_2(a))
}