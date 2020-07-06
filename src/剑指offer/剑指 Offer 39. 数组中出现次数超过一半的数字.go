package main

import "sort"

func majorityElement(nums []int) int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		}
		m[nums[i]] ++
	}

	max := 0
	res := 0
	for k, v := range m {
		if v > max {
			res = k
			max = v
		}
	}
	return res
}


func majorityElement_2(nums []int) int{

	sort.Ints(nums)

	return nums[len(nums) / 2]
}