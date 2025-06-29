package main

func mergeAlternately(word1 string, word2 string) string {
	mergedLen := len(word1) + len(word2)
	merged := make([]byte, 0, mergedLen)
	i := 0
	var long []byte
	var short []byte
	if len(word1) > len(word2) {
		long = []byte(word1)
		short = []byte(word2)
	} else {
		long = []byte(word2)
		short = []byte(word1)
	}
	for i < len(short) {
		merged = append(merged, word1[i])
		merged = append(merged, word2[i])
		i++
	}
	if len(word1) != len(word2) {
		merged = append(merged, long[i:]...)
	}
	return string(merged)
}
