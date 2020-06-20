package main

import "fmt"

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	mid := left
	if len(numbers) == 1 {
		return numbers[0]
	}
	if len(numbers) == 0 {
		return -1
	}
	for left<right{
		mid = (left+right) / 2
		if numbers[left] < numbers[right]{
			return numbers[left]
		}
		if numbers[mid]>numbers[right]{
			left = mid + 1
		}else if numbers[mid]<numbers[right]{
			right = mid
		}else{
			right--
		} // 去重
	}
	return numbers[left]
}

func main() {
	a := []int{3 ,3, 1, 3}
	fmt.Println(minArray(a))
}