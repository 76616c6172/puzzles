package main

import (
	"errors"
	"fmt"
	"strings"
)

func findCommonChar(s1 string, s2 string, s3 string) (rune, error) {

	for _, char := range s1 {

		// Only comparing 2 strings
		if len(s3) < 2 {
			if strings.Count(s2, string(char)) > 0 {
				return char, nil
			}
		}

		// Comparing 3 strings
		if strings.Count(s2, string(char)) > 0 && strings.Count(s3, string(char)) > 0 {
			return char, nil
		}

	}
	// If no common character is found, return an empty string
	return ' ', errors.New("no common character found")
}

func getItemPriority(r rune) int {
	itemPriority := int(r) - 96

	if int(r) < 91 {
		itemPriority = int(r) - 38
	}

	return itemPriority
}

func solveDay3Part1(s string) int {
	rucksackList := strings.Split(s, "\n")

	var sum int
	for _, sack := range rucksackList {
		half := len(sack) / 2
		c1 := sack[:half]
		c2 := sack[half:]

		commonItem, err := findCommonChar(c1, c2, "")
		if err != nil {
			panic(err)
		}

		sum += getItemPriority(commonItem)
	}

	return sum
}

func solveDay3Part2(s string) int {
	rucksackList := strings.Split(s, "\n")

	var sum, counter int
	for i := range rucksackList {
		counter++

		// Every three rucksacks find the common item in the last three sacks including the current one
		if counter == 3 {
			counter = 0
			commonItem, err := findCommonChar(rucksackList[i], rucksackList[i-1], rucksackList[i-2])
			if err != nil {
				panic(err)
			}

			sum += getItemPriority(commonItem)
		}

	}

	return sum
}

func solve_day3() {
	filepath := "puzzle-inputs/day3-input"
	input := get_input(filepath)
	solution1 := solveDay3Part1(input)
	solution2 := solveDay3Part2(input)
	fmt.Println("Day 3: Rucksack Reorganization = ", solution1)
	fmt.Println("Day 3: Part Two =", solution2)
}
