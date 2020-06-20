package main

import "fmt"

func main() {
	a := new(int)
	*a = 123
	fmt.Println(*a)
}
