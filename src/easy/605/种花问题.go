package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	res := false
	sum := 0
	temp := make([]int, len(flowerbed) + 2)
	copy(temp[1:], flowerbed)
	fmt.Println(temp)
	for i := 1; i < len(temp) -1; i ++ {
		if temp[i - 1] != 1 && temp[i+1] != 1 && temp[i] != 1{
			temp[i] =1
			sum += 1
 		}

	}
	if sum >= n {
		res = true
	}
	return res
}

func main()  {
	a := []int{1,0,0,0,1}

	fmt.Println(canPlaceFlowers(a, 1))
}