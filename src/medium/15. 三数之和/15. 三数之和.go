package main

import "sort"

func threeSum(nums []int) [][]int {
	a := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			for k := j+1;  k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					a = append(a, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return a
}

func threeSum_sort(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] >= 0 {
			break
		}
		left := i+1
		right := len(nums)-1
		for left  < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}

		for i < len(nums)-3 && nums[i] == nums[i+1]{
			i++
		}
	}


	return res
}