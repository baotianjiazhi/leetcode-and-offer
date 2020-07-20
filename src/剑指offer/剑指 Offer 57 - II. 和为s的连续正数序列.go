package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	res := [][]int{}
	n := []int{}
	for i := 1; i < target; i++ {
		n = append(n, i)
	}

	for j := 0; j < len(n); j++ {
		tmp := []int{}
		sum := 0
		tmp = append(tmp, n[j])
		sum += n[j]
		for m := j+1; m < len(n); m++ {
			sum += n[m]
			tmp = append(tmp, n[m])
			if sum > target {
				break
			} else if sum == target {
				res = append(res, tmp)
				break
			}
		}
	}

	return res
}


func main() {
	fmt.Println(findContinuousSequence(9))
}