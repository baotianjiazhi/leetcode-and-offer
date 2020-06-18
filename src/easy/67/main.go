package main

import (
	"fmt"
	"strconv"
)

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/
func addBinary(a string, b string) string {
	var l int
	var ret int = 0
	if len(b) == 0 && len(a) == 0 {
		return ""
	} else if len(a) == 0 {
		return b
	} else if len(b) == 0{
		return a
	}

	if len(a) > len(b){
		l = len(a)
	} else {
		l = len(b)
	}

	a, b = insert_0(a, b)
	fmt.Println("a:", a, "b:", b)
	temp := make([]byte, l+1, l+1)
	for i := l-1; i >= 0;i--{
		fmt.Println(int(a[i]))
		a_int, _ := strconv.Atoi(string(a[i]))
		b_int, _ := strconv.Atoi(string(b[i]))
		if (a_int + b_int + ret) == 0 {
			ret = 0
			temp[i+1] = 48
		} else if (a_int + b_int + ret) == 1{
			ret = 0
			temp[i+1] = 49
		} else if (a_int + b_int + ret) == 2{
			ret = 1
			temp[i+1] = 48
		} else if (a_int + b_int + ret) == 3 {
			ret = 1
			temp[i+1] = 49
		}
	}

	if ret == 1{
		temp[0] = 49
	}

	if temp[0] == 0 {
		temp = temp[1:]
	}
	fmt.Println(temp)
	return string(temp)
}

func insert_0(a string, b string) (string, string) {
	if len(a) > len(b) {
		b_ex := make([]byte, len(a), len(a))
		for i := range b {
			b_ex[len(b_ex)+i-len(b)] = b[i]
		}
		fmt.Println("b_ex", b_ex)
		return a, string(b_ex)
	} else if len(a) < len(b) {
		a_ex := make([]byte, len(b), len(b))
		for i := range a {
			a_ex[len(a_ex)+i-len(a)] = a[i]
			fmt.Println(a_ex)
		}
		fmt.Println("a_ex", a_ex)
		return string(a_ex), b
	}
	return a, b
}


func main() {
	a := "100"
	b := "110010"
	c := addBinary(a, b)
	fmt.Println(c)
}