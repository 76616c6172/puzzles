package main

import (
	"testing"
)

func Test_solve_day2(t *testing.T) {
	path := "puzzle-inputs/day2-example1"
	input := get_input(path)
	t.Run("Calculating score for part 1", func(t *testing.T) {
		got := calculateScore(input)
		want := 15
		if got != want {
			t.Errorf("\ninput:\n%v\n got: %v\nwant: %v", input, got, want)
		}
	})
	t.Run("Calculating score for part 2", func(t *testing.T) {
		got := calculateScore_part2(input)
		want := 12
		if got != want {
			t.Errorf("\ninput:\n%v\n got: %v\nwant: %v", input, got, want)
		}
	})
}
