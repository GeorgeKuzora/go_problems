package main

import (
	"reflect"
	"testing"
)

type testCasePlusOne struct {
	digits []int
	expected []int
}

func Test_plus_one(t *testing.T) {
	testCases := []testCasePlusOne {
		testCasePlusOne{
			digits: []int {1,2,3},
			expected: []int {1,2,4},
		},
		testCasePlusOne{
			digits: []int {1,2,9},
			expected: []int {1,1,2,0},
		},
		testCasePlusOne{
			digits: []int {9},
			expected: []int {1,0},
		},
		testCasePlusOne{
			digits: []int {0},
			expected: []int {1},
		},
		testCasePlusOne{
			digits: []int {9,9},
			expected: []int {1,0,0},
		},
	}
	for _, testCase := range testCases {
		actual := plusOne(testCase.digits)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Error("Failed")
			t.Error(actual)
			t.Error(testCase.expected)
		}
	}
}
