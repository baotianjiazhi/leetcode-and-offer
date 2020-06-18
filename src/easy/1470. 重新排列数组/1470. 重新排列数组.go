package main

func shuffle(nums []int, n int) []int {
	a := nums[:n]
	b := nums[n:]
	res := []int{}
	for i := 0;i < n; i++ {
		res = append(res, a[i])
		res = append(res, b[i])
	}

	return res
}

