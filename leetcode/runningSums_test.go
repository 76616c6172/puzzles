package main

import (
	"reflect"
	"testing"
)

// TestRunningSums is a unit test for the function runningSums
func TestRunningSums(t *testing.T) {
	testCases := [][]int{
		[]int{
			1, 2, 3, 4,
		},
		[]int{
			1, 1, 1, 1, 1,
		},
	}
	expectedOutputs := [][]int{
		[]int{
			1, 3, 6, 10,
		},
		[]int{
			1, 2, 3, 4, 5,
		},
	}

	for i, testCase := range testCases {
		got := runningSums(testCase)
		want := expectedOutputs[i]
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ninput: %v\n got: %v\nwant: %v", testCase, got, want)
		}
	}
}
