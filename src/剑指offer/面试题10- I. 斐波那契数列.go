package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	res := []int{}
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 1)
	for i := 3; i <= n; i++ {
		res = append(res, (res[i-1]+res[i-2]) %1000000007)
	}
	return res[len(res)-1]
}

func main()  {
	a := 3
	fmt.Println(fib(a))
}

