package main


var memory [][]bool

func isMatch(s string, p string) bool {
	memory = make([][]bool, len(s)+1)
	for i := 0; i < len(memory); i++ {
		memory[i] = make([]bool, len(p)+1)
	}

	return dp(0, 0, s, p)
}

func dp(i int, j int, s string, p string) bool {
	if memory[i][j] {
		return memory[i][j]
	}

	if len(p) == 0 {
		return len(s) == 0
	}

	ok := len(s) != 0 && (s[0] == p[0] || p[0] == '.') // 如果相等，或者为.匹配成功
	var res bool
	if len(p) >= 2 && p[1] == '*' {
		res = dp(i, j+2, s, p[2:]) || (ok && dp(i+1, j, s[1:], p))
	} else {
		res = ok && dp(i+1, j+1, s[1:], p[1:])
	}

	memory[i][j] = res
	return res


}
