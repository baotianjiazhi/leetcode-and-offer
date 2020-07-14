package main

func missingNumber(nums []int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := (i+j)>>1
		if m == nums[m] {
			i = m+1
		} else {
			j = m-1
		}
	}

	return i
}
