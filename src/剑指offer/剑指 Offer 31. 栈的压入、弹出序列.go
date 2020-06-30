package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, v := range pushed{
		stack = append(stack, v)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 3, 1, 2, 5}
	fmt.Println(validateStackSequences(a, b))
}