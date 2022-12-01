package main

import (
	"testing"
)

func Test_solve_day1(t *testing.T) {
	path := "./puzzle-inputs/day-1-example-1"
	testCase := getPuzzleInputFromFile(path)

	t.Run("Finding the highest calorie elf and adding up his calories", func(t *testing.T) {
		want := 24000
		got := getMostCarriedCalories(testCase)

		if got != want {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	})

	t.Run("Getting the top 3 calorie elves and adding up their calories", func(t *testing.T) {
		want := 45000
		testCase := getPuzzleInputFromFile(path)
		got := getSumOfTopThreeCalorieCarriers(testCase)

		if got != want {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	})
}
