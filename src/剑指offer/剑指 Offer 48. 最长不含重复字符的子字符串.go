package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(s); i++ {
		max := 0
		m := map [byte]int{}
		m[s[i]]++
		max++
		for j := i+1; j < len(s); j++ {
			if m[s[j]] != 0 {
				break
			}
			max++
			m[s[j]]++
		}
		if max > res {
			res = max
		}
	}

	return res
}


func main() {

	a := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(a))
}