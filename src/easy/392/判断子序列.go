package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	s1 := 0
	s2 := 0
	for s1 < len(s) && s2 < len(t) {
		if s[s1] == t[s2] {
			s1++
		}
		s2++
	}

	if s1 == len(s) {
		return true
	}

	return false
}
