package main

import (
	"fmt"
	"strings"
)

func swap(a, b *int) {
	s := []string{"dsfsafa", "dasfasdf"}
	fmt.Println(strings.Join(s, ""))
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