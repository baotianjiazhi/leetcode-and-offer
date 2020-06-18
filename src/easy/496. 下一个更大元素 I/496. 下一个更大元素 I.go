package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	a := map[int]int{}
	for k, v := range nums2{
		a[v] = k
	}
	for i:= 0; i < len(nums1); i++ {
		res[i] = -1
		for j := a[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res[i] = nums2[j]
				break
			}
		}

	}

	return res
}


func main() {
	a := []int{4, 1, 2}
	b := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(a, b))
}
