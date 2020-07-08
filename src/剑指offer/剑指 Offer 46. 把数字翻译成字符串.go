package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	n := strconv.Itoa(num)
	if len(n) < 2{
		return 1
	}
	dp := []int{1, 1}
	for i := 2; i <= len(n); i++ {
		dp = append(dp, 0)
		dp[i] = dp[i-1]
		x, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(n[i-2]), string(n[i-1])))
		fmt.Println(x)
		if x <= 25 && x >= 10{
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	fmt.Println(dp)
	return dp[len(n)]
}


func main() {
	fmt.Println(translateNum(1068385902))
}