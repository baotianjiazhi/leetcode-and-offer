package main

import (
	"fmt"
)

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
func main() {
	x := []int{1, 2}
	fmt.Println(x[1:])
}
