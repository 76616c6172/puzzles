package main

import (
	"testing"
)

func Test_solve_day1(t *testing.T) {
	path := "./puzzle-inputs/day-1-example-1"
	testcase := get_input(path)

	t.Run("Finding highest calorie elf & adding up calories", func(t *testing.T) {
		want := 24000
		got := getHighestCalories(testcase)

		if got != want {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	})

	t.Run("Getting  top 3 calorie elves & adding up calories", func(t *testing.T) {
		want := 45000
		testcase := get_input(path)
		got := addTop3Calories(testcase)

		if got != want {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	})
}
