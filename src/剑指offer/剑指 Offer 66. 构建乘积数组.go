package main

func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	num1 := make([]int, len(a))
	num1[0] = 1
	num2 := make([]int, len(a))
	num2[len(a)-1] = 1
	res := []int{}
	for i := 1;i < len(a); i++ {
		num1[i] = num1[i-1] * a[i-1]
	}

	for i := len(a)-2; i >= 0; i-- {
		num2[i] = num2[i+1] * a[i+1]
	}

	for i := 0; i < len(a); i++ {
		res = append(res, num1[i]*num2[i])
	}
	return res
}
