package main

func isAnagram(s string, t string) bool {
    sLen := len(s)
    tLen := len(t)
    if sLen != tLen {
    	return false
    }

    sSet := make(map[rune]int, sLen)
    for _, v := range s {
    	sSet[v]++
    }
    for _, v := range t {
    	r, ok := sSet[v]
    	if ok == false {
	   		return false
     	}
		r--
		sSet[v] = r
    }
    for _, v := range s {
    	r := sSet[v]
	    if r != 0 {
			return false
		}
    }
    return true
}
