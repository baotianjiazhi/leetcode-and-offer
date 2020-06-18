package main

func rob(nums []int) int {
	f := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	f[0] = nums[0]
	f[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i ++ {
		f[i] = max(f[i-1], f[i-2] + nums[i])
	}

	return f[len(nums) -1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}