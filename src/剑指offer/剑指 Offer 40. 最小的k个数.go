package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	quicksort(arr, 0, len(arr)-1)
	return arr[:k]
}

func quicksort(arr []int, left, right int) {

	if left < right {
		i := left
		j := right
		tmp := arr[i]
		for i < j {
			for i < j && arr[j] > tmp{
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}

			for i < j && arr[i] < tmp {
				arr[j] = arr[i]
				i++
			}

			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = tmp
		quicksort(arr, left, i-1)
		quicksort(arr, i+1, right)
	}
}

func main() {

	a := []int{1, 2, 5, 7, 4, 2, 4}

	fmt.Println(getLeastNumbers(a, 4))
}