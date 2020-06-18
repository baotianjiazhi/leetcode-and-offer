package main

import "fmt"

func main() {
	a := []int{1, 2}
	fmt.Println(a[:len(a)-1])
}
