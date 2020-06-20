package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[right] > nums[mid] {
			if target > nums[mid] && target <= nums[right]  {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			}  else {
				left = mid + 1
			}
		}
	}

	return -1
}

func main() {
	a := []int{4,5,6,7,0,1,2}
	fmt.Println(search(a, 0))
}