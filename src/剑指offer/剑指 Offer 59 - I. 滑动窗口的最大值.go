package main


func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := []int{}

	for i := 0; i+k-1 < len(nums); i++ {
		max := nums[i]
		for  x := i; x < i+k; x++ {
			if max < nums[x] {
				max = nums[x]
			}
		}
		res = append(res, max)
	}

	return res
}
