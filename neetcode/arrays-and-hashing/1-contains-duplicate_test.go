package neetcode

import (
	"testing"
)

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Input: nums = [1,2,3,1]
Output: true

Input: nums = [1,2,3,4]
Output: false

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func Test_Contains_duplicate(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	solutions := []bool{
		true,
		false,
		true,
	}

	for i, input := range testCases {
		got := Contains_duplicate(input)
		want := solutions[i]

		if got != want {
			t.Errorf("\ninput: %v\n got: %v\nwant: %v", input, got, want)
			return
		}

	}
}
