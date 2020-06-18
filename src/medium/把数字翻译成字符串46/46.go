package main

import (
	"strconv"
)

func translateNum(num int) int {
	n := strconv.Itoa(num)
	p, q, r := 0, 0, 1

	for i := 0; i < len(n); i ++  {
		p = q
		q = r
		r = 0
		r += q
		if i == 0 {
			continue
		}

		pre := n[i-1:i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}

	return r
}
