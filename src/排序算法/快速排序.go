package main

import "fmt"

func quick_sort(list []int, left, right int) {
	if left < right {
		i := left
		j := right
		t := list[i]
		for i < j {
			for i<j && list[j] > t {
				j--
			}
			if i < j {
				list[i] = list[j]
				i++
			}
			for i < j && list[i] < t {
				i++
			}

			if i < j {
				list[j] = list[i]
				j--
			}
		}

		list[i] = t
		quick_sort(list, left, j-1)
		quick_sort(list, j+1, right)
	}
}


func main() {
	a := []int{3, 4, 1, 2, 4, 5, 8, 6}
	quick_sort(a, 0, len(a)-1)
	fmt.Println(a)
}