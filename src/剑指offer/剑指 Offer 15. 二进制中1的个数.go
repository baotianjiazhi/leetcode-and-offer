package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	var flag uint32 = 1
	for flag != 0 {

		if (num & flag != 0) {
			count ++
		}
		flag <<= 1
	}

	return count
}

func hammingWeight_1(num uint32) int {
	count := 0
	for num != 0 {
		num = (num -1) & num
		count++
	}

	return count
}

func main() {
	var n uint32 = 123
	fmt.Println(hammingWeight(n))
}
