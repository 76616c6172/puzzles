package main

import (
	"testing"
)

func Test_solve_day3(t *testing.T) {
	path := "./puzzle-inputs/day3-example1"
	testcase := get_input(path)

	t.Run("Testing day3", func(t *testing.T) {
		got := solveDay3Part1(testcase)
		want := 157

		if got != want {
			t.Errorf("\ninput: %v\n got: %v\nwant: %v", testcase, got, want)
		}
	})

	t.Run("Testing day3 part 2", func(t *testing.T) {
		got := solveDay3Part2(testcase)
		want := 70

		if got != want {
			t.Errorf("\ninput: %v\n got: %v\nwant: %v", testcase, got, want)
		}
	})
}
