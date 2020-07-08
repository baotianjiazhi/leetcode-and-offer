package main

import (
	"fmt"
	"strings"
)

func minNumber(nums []int) string {
	var res strings.Builder
	quicksort_1(nums, 0, len(nums)-1)
	for _, v := range nums{
		res.WriteString(fmt.Sprintf("%d", v))
	}

	return res.String()
}

func quicksort_1(nums []int, left, right int){

	if left < right {
		i := left
		j := right
		tmp := nums[i]
		for i < j {
			for i < j && !compareNumber(tmp, nums[j]) {
				j--
			}

			if i < j {
				nums[i] = nums[j]
				i++
			}

			for i < j && compareNumber(tmp, nums[i]) {
				i++
			}

			if i < j{
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = tmp
		quicksort_1(nums, left, i-1)
		quicksort_1(nums, i+1, right)
	}
}

func compareNumber(a, b int) bool {
	str_1 := fmt.Sprintf("%d%d", a, b)
	str_2 := fmt.Sprintf("%d%d", b, a)

	if str_1 > str_2 {
		return true
	} else {
		return false
	}
}

func main() {

	a := []int{10, 2, 3, 5, 7, 6}
	fmt.Println(minNumber(a))
}