package main


func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	rk, ans := -1, 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rk + 1 < len(s) && m[s[rk+1]] == 0{
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}

	return ans
}


func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}