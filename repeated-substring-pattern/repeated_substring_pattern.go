package main

func repeatedSubstringPattern(s string) bool {
	fullS := s + s
	testS := fullS[1:len(fullS)-1]
	for i := 0; i <= len(testS)-len(s); i++ {
		if s == testS[i:i+len(s)] {
			return true
		}
	}
	return	false
}
