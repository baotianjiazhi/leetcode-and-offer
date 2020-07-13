package main

import "fmt"

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
func main() {
	for i := 0; i < 1; i++ {
		for j := 0; j < 1; j++ {
			break
		}
		fmt.Println(1)
	}
}
func min(a, b, c int) (int) {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}