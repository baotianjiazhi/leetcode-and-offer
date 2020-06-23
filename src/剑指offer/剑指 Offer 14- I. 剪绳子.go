package main

/**
f(i) = f(i-j) * f(j)
f(x)表示长度为x绳子剪断之后的最大乘积和，但是在这个题中其实长度为1、2、3时不剪最好，但是要求必须给绳子必须剪，所以单独给
长度为1、2、3的绳子时，最优解是1、1、2，但是当长度大于3时，剪下来还剩1、2、3长度的绳子时，最优解为1、2、3
 */

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return  2
	}

	a := make([]int, n+1)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j ++ {
			if a[i-j] * a[j] > max {
				max = a[i-j] * a[j]
			}
		}
		a[i] = max
	}
	return a[len(a)-1]
}

