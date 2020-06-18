package main

import "fmt"

func waysToStep(n int) int {
	f := make([]int, n)
	f[0] = 1
	if n > 1 {
		f[1] = 2
	}
	if n > 2 {
		f[2] = 4
	}
	for i := 3; i < n; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3])%1000000007
	}

	return f[n-1]
}


func main() {
	fmt.Println(waysToStep(3))
}