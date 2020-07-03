package main

import "fmt"

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a[:2])

}
