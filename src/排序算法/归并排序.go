package main

func Mergesort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	mid := (l+r)>>1
	Mergesort(arr, l, mid)
	Mergesort(arr, mid+1, r)
	Merge(arr, l, mid, r)
}

func Merge(arr []int, l, mid, r int) {
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-1] = arr[i]
	}
	left := 1
	right := mid+1
	for i := l; i <= r; i++ {
		if left > mid {
			arr[i] = temp[right-1]
			right++
		} else if right > r {
			arr[i] = temp[left-1]
			left++
		} else if temp[left -1] < temp[right-1] {
			arr[i] = temp[right-1]
			right++
		} else {
			arr[i] = temp[left-1]
			left++
		}
	}
}
