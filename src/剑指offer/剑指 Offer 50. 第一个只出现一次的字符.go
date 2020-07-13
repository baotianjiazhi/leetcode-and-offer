package main

import "fmt"

func firstUniqChar(s string) byte {

	if len(s) == 0 {
		return 32
	}

	s_c := []byte(s)
	m := map [byte]int{}

	var res byte
	for i := 0; i < len(s_c); i++ {
		flag := 0
		if _, ok := m[s_c[i]]; ok {
			continue
		}
		m[s_c[i]]++
		for j := i+1; j < len(s_c); j++ {
			if s[j] == s[i] {
				flag = 1
			}
		}

		if flag == 0 {
			res = s[i]
			break
		}
	}

	if res == 0 {
		return 32
	}
	return res
}

func main() {
	s := "cc"
	fmt.Println(firstUniqChar(s))
}