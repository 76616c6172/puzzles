package neetcode

import (
	"testing"
)

/*
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
*/

func TestIsAnagram(t *testing.T) {
	testCases := [][2]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"a", "ab"},
	}
	solutions := []bool{
		true,
		false,
		false,
	}

	for i := range testCases {
		got := IsAnagram(testCases[i][0], testCases[i][1])
		want := solutions[i]

		if got != want {
			t.Errorf("%s got: %v, want: %v", testCases[i], got, want)
		}
	}
}
