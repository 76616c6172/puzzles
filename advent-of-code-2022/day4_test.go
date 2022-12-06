package main

import (
	"testing"
)

func Test_solve_day4(t *testing.T) {
	path := "./puzzle-inputs/day4-example1"
	testcase := get_input(path)

	t.Run("Testing day4", func(t *testing.T) {
		got := solveDay4Part1(testcase)
		want := 2

		if got != want {
			t.Errorf("\ninput:\n%v\n got: %v\nwant: %v", testcase, got, want)
		}
	})
}
