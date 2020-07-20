package main

import "strings"

func reverseLeftWords(s string, n int) string {
	res := []string{}
	res = append(res, s[n:])
	res = append(res, s[:n])
	return strings.Join(res, "")

}
