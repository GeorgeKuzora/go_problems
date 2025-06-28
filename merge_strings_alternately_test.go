package main

import "testing"

type testCase struct {
	word1 string
	word2 string
	expected string
}

func Test_mergeAlternately(t *testing.T)  {
	testCases := []testCase{
		testCase{
			word1: "a",
			word2: "b",
			expected: "ab",
		},
		testCase{
			word1: "abc",
			word2: "def",
			expected: "adbecf",
		},
		testCase{
			word1: "abc",
			word2: "defgh",
			expected: "adbecfgh",
		},
		testCase{
			word1: "abcde",
			word2: "fgh",
			expected: "afbgchde",
		},
	}	

	for _, testCase := range testCases {
		actual := mergeAlternately(testCase.word1, testCase.word2)
		if actual != testCase.expected {
			t.Errorf("Incorrect result: expected: %s, actual: %s", testCase.expected, actual)
		}
	}
}
