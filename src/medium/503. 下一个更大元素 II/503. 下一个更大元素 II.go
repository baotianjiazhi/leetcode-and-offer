package main

import "fmt"

type Tem struct {
	k int
	v int
}
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for  k := range res {
		res[k] = -1
	}
	stack := []Tem{}
	for i := 0; i < len(nums) * 2; i++ {
		n := i % len(nums)
		for len(stack) > 0 {
			if nums[n] > stack[len(stack)-1].v{
				res[stack[len(stack)-1].k] = nums[n]
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, Tem{k:n, v:nums[n]})
	}

	return res
}

func main() {
	a := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(a))
}