package main

import "testing"

type FtiotfoTestCase struct {
	haystack string
	needle string
	expected int
}

func Test_strStr(t *testing.T) {
	testCases := []FtiotfoTestCase{
		FtiotfoTestCase{
			haystack: "a",
			needle: "a",
			expected: 0,
		},
		FtiotfoTestCase{
			haystack: "abcd",
			needle: "abc",
			expected: 0,
		},
		FtiotfoTestCase{
			haystack: "abcd",
			needle: "cd",
			expected: 2,
		},
		FtiotfoTestCase{
			haystack: "abcd",
			needle: "ww",
			expected: -1,
		},
		FtiotfoTestCase{
			haystack: "abcd",
			needle: "wwww",
			expected: -1,
		},
		FtiotfoTestCase{
			haystack: "mississippi",
			needle: "pi",
			expected: 9,
		},
	}

	for _, testCase := range testCases {
		actual := strStr(testCase.haystack, testCase.needle)
		if actual != testCase.expected {
			t.Errorf("Failed: expected %d, actual %d", testCase.expected, actual)
		}
	}
}
