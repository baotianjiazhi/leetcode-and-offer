package main

import (
	"fmt"
	"sort"
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	if nums[0] * 2 > target {
		return 0
	}
	var res int64 = 0
	left := 0
	right := len(nums) -1
	for left <= right {
		if nums[left] + nums[right] <= target {
			res += myPow(2, right-left)
			left++
		} else {
			right--
		}
	}

	res = res % (1e9 + 7)
	return int(res)
}


func myPow(x int, n int) int64 {
	var res int = 1
	if n == 0 {
		return 1
	} else if n == 1 {
		return int64(x)
	}

	for n != 0 {
		if n & 1 != 0 {
			res = res * x
		}

		x = x * x
		n >>= 1
	}

	return int64(res)
}

func main() {
	//nums := []int{7, 10, 7, 5, 6, 7, 3,4 ,9, 6}
	fmt.Println(myPow(2, 1))
}