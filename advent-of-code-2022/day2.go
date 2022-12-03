package main

import (
	"fmt"
	"strings"
)

// Encode the byte values from the script of the game
const (
	MY_ROCK        byte = 'X'
	MY_PAPER       byte = 'Z'
	MY_SCISSORS    byte = 'Y'
	THEIR_ROCK     byte = 'A'
	THEIR_PAPER    byte = 'B'
	THEIR_SCISSORS byte = 'C'
)

const (
	POINTS_FOR_MY_ROCK     = 1
	POINTS_FOR_MY_PAPER    = 2
	POINTS_FOR_MY_SCISSORS = 3
)

// Define the points awarded for the outcome of the game
const (
	LOSS = 0
	DRAW = 3
	WIN  = 6
)

var outcomeScore = map[byte]int{
	'X': LOSS,
	'Y': DRAW,
	'Z': WIN,
}

var pointsGivenMyShape = map[byte]int{
	'R': 1, // rock
	'P': 2, // paper
	'S': 3, // scissors
}

// Encodes the rules of the game, rock beats scissors and so on..
// RP[0] represents my opponents choice of rock
// RP[1] represents my choice of paper
var pointsGivenForOutcome = map[string]int{
	"RP": WIN,
	"PS": WIN,
	"SR": WIN,
	"PR": LOSS,
	"SP": LOSS,
	"RS": LOSS,
	"SS": DRAW,
	"RR": DRAW,
	"PP": DRAW,
}

// re_encode_input_data replaces the stupid encoding (A and X are both rock) with a sensible encoding (R means rock) that doesn't vary depending on wether it's me or my oponent playing the option
func re_encode_input_data(s string) string {
	s = strings.ReplaceAll(s, "A", "R")
	s = strings.ReplaceAll(s, "B", "P")
	s = strings.ReplaceAll(s, "C", "S")

	s = strings.ReplaceAll(s, "X", "R")
	s = strings.ReplaceAll(s, "Y", "P")
	s = strings.ReplaceAll(s, "Z", "S")

	return s
}

func re_encode_only_their_shape_selection(s string) string {
	s = strings.ReplaceAll(s, "A", "R")
	s = strings.ReplaceAll(s, "B", "P")
	s = strings.ReplaceAll(s, "C", "S")

	return s
}

func getScoreFromOutCome(s string) int {
	var score int
	theirInput := (s[0])
	myInput := (s[2])
	key := string(theirInput) + string(myInput)
	score += pointsGivenForOutcome[key]

	return score
}

func getScoreFromMyShape(s string) int {
	var score int
	myInput := (s[2])
	score += pointsGivenMyShape[myInput]

	return score
}

func calculateScore(s string) int {
	sensiblyEncodedData := re_encode_input_data(s)
	gamesequence := strings.Split(sensiblyEncodedData, "\n")

	var score int
	for _, game := range gamesequence {
		score += getScoreFromMyShape(game) + getScoreFromOutCome(game)
	}

	return score
}

func pointsFromMyShapeSelection(g string, n int) int {
	theirShape := g[0]

	switch n {
	case LOSS:

		if theirShape == 'R' {
			return POINTS_FOR_MY_SCISSORS
		} else if theirShape == 'P' {
			return POINTS_FOR_MY_ROCK
		} else {
			return POINTS_FOR_MY_PAPER
		}

	case DRAW:

		if theirShape == 'R' {
			return POINTS_FOR_MY_ROCK
		} else if theirShape == 'P' {
			return POINTS_FOR_MY_PAPER
		} else {
			return POINTS_FOR_MY_SCISSORS
		}

	case WIN:

		if theirShape == 'R' {
			return POINTS_FOR_MY_PAPER
		} else if theirShape == 'P' {
			return POINTS_FOR_MY_SCISSORS
		} else {
			return POINTS_FOR_MY_ROCK
		}

	}

	return 0
}

func calculateScore_part2(s string) int {
	data := re_encode_only_their_shape_selection(s)
	gamesequence := strings.Split(data, "\n")

	var score int
	for _, game := range gamesequence {

		outcome := outcomeScore[byte(game[2])]
		score += outcome + pointsFromMyShapeSelection(game, outcome)
	}

	return score
}

func solve_day2() {
	filepath := "puzzle-inputs/day2-part1"
	input := get_input(filepath)
	score := calculateScore(input)
	score2 := calculateScore_part2(input)
	fmt.Println("Day 2: Rock Paper Scissors =", score)
	fmt.Println("Day 2: Part Two =", score2)
}
