package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		for j := i+1; j < len(nums)-2; j++ {
			left := j+1
			right := len(nums)-1
			for left < right {
				sum := nums[left] + nums[right] + nums[i] + nums[j]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for right > left && nums[right] == nums[right-1]{
						right--
					}
					left++
					right--
				}
			}
			for j < len(nums) - 2 && nums[j] == nums[j+1]{
				j++
			}
		}

		for i < len(nums) - 3 && nums[i] == nums[i+1]{
			i++
		}
	}

	return res
}