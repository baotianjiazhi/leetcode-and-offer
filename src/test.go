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
	a := "120"
	b := "210"
	fmt.Println(a > b)
}
