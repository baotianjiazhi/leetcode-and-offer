package main

import "fmt"

func searchRange(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	res := []int{-1,-1}
	for ;i < len(nums); i++ {
		if nums[i] <=
	}

	if res[0] != -1 && res[1] == -1 {
		res[1] = res[0]
	}
	return res
}

func main()  {
	a := []int{1, 3, 3, 4, 3}
	fmt.Println(searchRange(a, 3))
}