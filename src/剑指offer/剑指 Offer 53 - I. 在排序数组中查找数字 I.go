package main

func search(nums []int, target int) int {
	i := 0
	j := len(nums)-1

	for i <= j {
		m := (i+j)>>1
		if nums[m] <= target {
			i = m+1
		} else {
			j = m-1
		}
	}

	right := i

	i = 0
	j = len(nums)-1

	for i <= j {
		m := (i+j)>>1
		if nums[m] < target {
			i = m+1
		} else {
			j = m-1
		}
	}

	left := j
	return right - left - 1
}
