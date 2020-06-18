package main

func numPairsDivisibleBy60(time []int) int {
	m := make([]int, 60)
	cnt := 0
	for _, n := range time {
		if n % 60 == 0 {
			cnt += m[0]
		} else {
			cnt += m[60-n%60]
		}
		m[n%60]++
	}
	return cnt
}
