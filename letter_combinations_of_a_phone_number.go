package main

import "fmt"

func main() {
fmt.Println(letterCombinations("hello world"))
}

func letterCombinations(digits string) []string {
	phone_map := map[string][]string {
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	answer := []string{digits}
	return answer
}

