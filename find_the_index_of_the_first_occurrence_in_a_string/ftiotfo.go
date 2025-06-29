package main

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
    stackLen := len(haystack)
    if needleLen > stackLen {
    	return -1
    }
    needleSum := 0
    for _, v := range needle {
	   	needleSum = needleSum + int(v)
    }

    for i := 0; i <= stackLen - needleLen; i++ {
	   	stackSum := 0
	   	for j := i; j < i+needleLen; j++ {
			if haystack[j] == needle[j-i] {
				stackSum = stackSum + int(haystack[j])
			}
		}
		if needleSum == stackSum {
			return i
		}
    }
    return -1
}
