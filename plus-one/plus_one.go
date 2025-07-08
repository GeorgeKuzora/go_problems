package main

func plusOne(digits []int) []int {
    lenD := len(digits)

    if digits[lenD - 1] < 9 {
	   	digits[lenD - 1] = digits[lenD - 1] + 1
		return digits
    }

    newD := make([]int, 1, lenD + 1)
    newD[0] = 1
    digits[lenD - 1] = 0
    newD = append(newD, digits...)
    return newD
}
