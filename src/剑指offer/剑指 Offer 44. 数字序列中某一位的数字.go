package main

import "strconv"

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = 9 * digit * start
	}

	num := start + (n-1) / digit
	num_str := strconv.Itoa(num)
	return int(num_str[(n-1) % digit] - '0')
}
