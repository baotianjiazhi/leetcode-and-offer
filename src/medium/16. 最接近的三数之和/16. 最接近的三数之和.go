package main

import (
	"sort"
	"fmt"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		left := i+1
		right := len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
			if Abs(res, target) > Abs(sum, target) {
				res = sum
			}
		}
	}

	return res
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	a := []int{-100,-98,-2,-1}
	fmt.Println(threeSumClosest(a, -101))
}