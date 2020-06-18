package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	sum_blank := 0
	temp := []byte(s)
	for _, v := range temp {
		if v == 32 {
			sum_blank += 1
		}
	}
}
func replaceSpace_1(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func main() {
	fmt.Println(replaceSpace("123"))
}