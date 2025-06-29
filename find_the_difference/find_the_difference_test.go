package main

import "testing"

type findTheDifferenceTestCase struct {
	s string
	t string
	expected string
}

func Test_findTheDifference(t *testing.T) {
	testCases := []findTheDifferenceTestCase{
		findTheDifferenceTestCase{
			s: "",
			t: "a",
			expected: "a",
		},
		findTheDifferenceTestCase{
			s: "abc",
			t: "byca",
			expected: "y",
		},
		findTheDifferenceTestCase{
			s: "abc",
			t: "abcy",
			expected: "y",
		},
		findTheDifferenceTestCase{
			s: "abc",
			t: "abca",
			expected: "a",
		},
	}
	for _, testCase := range testCases {
		actual := findTheDifference(testCase.s, testCase.t)
		actualS := string(actual)
		if actualS != testCase.expected {
			t.Errorf("Failed: expected: %s, actual %s", testCase.expected, actualS)
		}
	}
}
