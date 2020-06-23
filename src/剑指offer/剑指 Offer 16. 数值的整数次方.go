package main

func myPow(x float64, n int) float64 {
	var res float64 = 1
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1/x
		n = -n
	}

	for n != 0 {
		if n & 1 != 0 {
			res = res * x
		}

		x *= x
		n >>= 1
	}

	return res
}
