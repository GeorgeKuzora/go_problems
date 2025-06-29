package main

import "testing"

type validAnagramTestCase struct {
	s string
	t string
	expected bool
}

func Test_validAnagram(t *testing.T) {
	testCases := []validAnagramTestCase {
		validAnagramTestCase{
			s: "a",
			t: "a",
			expected: true,
		},
		validAnagramTestCase{
			s: "abc",
			t: "bca",
			expected: true,
		},
		validAnagramTestCase{
			s: "abca",
			t: "baca",
			expected: true,
		},
		validAnagramTestCase{
			s: "a",
			t: "b",
			expected: false,
		},
		validAnagramTestCase{
			s: "a",
			t: "ab",
			expected: false,
		},
		validAnagramTestCase{
			s: "ab",
			t: "ca",
			expected: false,
		},
		validAnagramTestCase{
			s: "abca",
			t: "cbac",
			expected: false,
		},
	}
	for _, testCase := range testCases {
		actual := isAnagram(testCase.s, testCase.t)
		if actual != testCase.expected {
			t.Errorf("Case %s-%s failed: expected: %t, actual: %t", testCase.s, testCase.t, testCase.expected, actual)
		}
	}
}
