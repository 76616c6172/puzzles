package main

import "testing"

func TestGetPivotIndex(t *testing.T) {
	testCases := [][]int{
		[]int{1, 7, 3, 6, 5, 6},
		[]int{1, 2, 3},
		[]int{2, 1, -1},
	}
	solUtions := []int{
		3,
		-1,
		0,
	}

	for i, testCase := range testCases {
		got := getPivotIndex(testCase)
		want := solUtions[i]
		if got != want {
			t.Errorf("\ninput: %v\n got: %v\nwant: %v", testCase, got, want)
		}
	}
}
