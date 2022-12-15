package main

import (
	"testing"
)

func Test_solve_day4(t *testing.T) {
	path := "./puzzle-inputs/day4-example1"
	testcase := get_input(path)
	t.Run("Testing day4 part 1", func(t *testing.T) {
		got := solveDay4Part1(testcase)
		want := 2
		if got != want {
			t.Errorf("\ninput:\n%v\n got: %v\nwant: %v", testcase, got, want)
		}
	})
	t.Run("Testing day4 part 2", func(t *testing.T) {
		got := solveDay4Part2(testcase)
		want := 4
		if got != want {
			t.Errorf("\ninput:\n%v\n got: %v\nwant: %v", testcase, got, want)
		}
	})
}
