package main

import "testing"

type testCaseRepeatedSubstringPattern struct {
	s string
	expected bool
}

func Test_repeatedSubstringPattern(t *testing.T) {
	testCases := []testCaseRepeatedSubstringPattern {
		testCaseRepeatedSubstringPattern{
			s: "a",
			expected: false,
		},
		testCaseRepeatedSubstringPattern{
			s: "aa",
			expected: true,
		},
		testCaseRepeatedSubstringPattern{
			s: "abab",
			expected: true,
		},
		testCaseRepeatedSubstringPattern{
			s: "abcabd",
			expected: false,
		},
		testCaseRepeatedSubstringPattern{
			s: "aba",
			expected: false,
		},
		testCaseRepeatedSubstringPattern{
			s: "aaa",
			expected: true,
		},
		testCaseRepeatedSubstringPattern{
			s: "aaaab",
			expected: false,
		},
		testCaseRepeatedSubstringPattern{
			s: "ababab",
			expected: true,
		},
	}
	for _, testCase := range testCases {
		actual := repeatedSubstringPattern(testCase.s)
		if actual != testCase.expected {
			t.Errorf("Failed case: %s. Expected: %t. Actual: %t", testCase.s, testCase.expected, actual)
		}
	}
}
