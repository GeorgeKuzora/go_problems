package main

func findTheDifference(s string, t string) byte {
	original := []byte(s)
	shuffled := []byte(t)

	quickSort(original)
	quickSort(shuffled)

	for i := 0; i < len(original); i++ {
		if original[i] != shuffled[i] {
			return shuffled[i]
		}
	}
	return shuffled[len(shuffled) - 1]
}

func quickSort(s []byte) {
	if len(s) <= 1 {
		return
	}
	low := 0
	high := len(s) - 1
	pivotId := high
	pivotVal := s[pivotId]

	futurePivotId := -1

	for i:=low; i<=high; i++ {
		if s[i] < pivotVal {
			futurePivotId++
			temp := s[i]
			s[i] = s[futurePivotId]
			s[futurePivotId] = temp
		}
	}
	futurePivotId++
	temp := s[futurePivotId]
	s[futurePivotId] = s[pivotId]
	s[pivotId] = temp

	quickSort(s[:futurePivotId])
	quickSort(s[futurePivotId+1:])
}
