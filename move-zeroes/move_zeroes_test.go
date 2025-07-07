package main

import (
	"testing"
)

type testCaseMoveZeros struct {
	nums []int
	expected []int
	id string
}

func Test_moveZeroes(t *testing.T)  {
	testCases := []testCaseMoveZeros{
		testCaseMoveZeros {
			nums: []int{0,},
			expected: []int {0,},
			id: "1",
		},
		testCaseMoveZeros {
			nums: []int{1,},
			expected: []int {1,},
			id: "2",
		},
		testCaseMoveZeros {
			nums: []int{0,1},
			expected: []int {1,0},
			id: "3",
		},
		testCaseMoveZeros {
			nums: []int{0,0,1},
			expected: []int {1,0,0},
			id: "4",
		},
		testCaseMoveZeros {
			nums: []int{0,2,0,1},
			expected: []int {2,1,0,0},
			id: "5",
		},
		testCaseMoveZeros {
			nums: []int{0,1,0,2},
			expected: []int {1,2,0,0},
			id: "6",
		},
		testCaseMoveZeros {
			nums: []int{0,1,2,0},
			expected: []int {1,2,0,0},
			id: "7",
		},
	}
	for _, testCase := range testCases {
		moveZeroes(testCase.nums)
		for i, num := range testCase.nums {
			if num != testCase.expected[i] {
				t.Errorf("Failed ID %s", testCase.id)
				t.Error(testCase.nums)
				t.Error(testCase.expected)
			}
		}
	}
}
