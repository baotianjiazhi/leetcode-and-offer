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
	p1 := len(s)-1
	temp := []byte(s)
	for _, v := range temp {
		if v == 32 {
			sum_blank += 1
		}
	}

	for i := 0; i < sum_blank; i++ {
		temp = append(temp, 32, 32)
	}
	p2 := len(temp)-1
	for p1 != p2 {
		if temp[p1] == 32 {
			temp[p2] = 48
			p2--
			temp[p2] = 50
			p2--
			temp[p2] = 37
			p2--
			p1--
		} else if temp[p1] != 32 {
			temp[p2] = temp[p1]
			p1--
			p2--
		}

	}

	return string(temp)
}

func replaceSpace_1(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func main() {
	fmt.Println(replaceSpace("1 2 3"))
}