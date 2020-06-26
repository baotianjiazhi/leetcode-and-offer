package main

func exchange(nums []int) []int {

	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	i := 0
	j := len(nums) - 1

	for i != j {
		for nums[i] & 1 != 0 && i < j{
			i++
		}

		for nums[j] & 1 == 0 && i < j {
			j--
		}

		swap(nums, i, j)
	}

	return nums
}

func swap(nums []int, i, j int) {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
}