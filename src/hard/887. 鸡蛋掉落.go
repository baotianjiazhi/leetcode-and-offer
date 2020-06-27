package main

type Building struct {
	Numbers int
	Floor int
}

var m map[Building]int

func superEggDrop(K int, N int) int {

	m = map[Building]int{}
	return dp(K, N)
}

func dp(K, N int) int {
	if K == 1 {
		return N
	}

	if K == 0 {
		return 0
	}



	if value, ok := m[Building{
		Numbers: K,
		Floor: N,
	}]; ok {
		return value
	}

	res := N
	for i := 1; i < N+1; i++ {
		res = min(res, max(dp(K, N-i), dp(K-1, i-1))+1)
	}
	m[Building{Numbers: K, Floor: N}] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}


func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}