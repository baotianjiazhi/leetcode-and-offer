package main

import "fmt"

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
