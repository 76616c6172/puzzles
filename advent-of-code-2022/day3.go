package main

import (
	"errors"
	"fmt"
	"strings"
)

func findCommonChar(s1 string, s2 string, s3 string) (rune, error) {
	for _, char := range s1 {

		// If comparing 2 strings
		if len(s3) < 2 {
			if strings.Count(s2, string(char)) > 0 {
				return char, nil
			}
		}

		// if comparing 3 strings
		if strings.Count(s2, string(char)) > 0 && strings.Count(s3, string(char)) > 0 {
			return char, nil
		}
	}

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
	sackList := strings.Split(s, "\n")

	var itemPrioritySum int
	for _, sack := range sackList {
		middleIndex := len(sack) / 2

		compartment1 := sack[:middleIndex]
		compartment2 := sack[middleIndex:]

		commonItem, err := findCommonChar(compartment1, compartment2, "")
		if err != nil {
			panic(err)
		}

		itemPrioritySum += getItemPriority(commonItem)
	}

	return itemPrioritySum
}

func solveDay3Part2(s string) int {
	sackList := strings.Split(s, "\n")

	var itemPrioritySum, sackCounter int
	for i := range sackList {
		sackCounter++

		// Every three rucksacks find the common item in the last three sacks including the current one
		if sackCounter == 3 {
			sackCounter = 0
			commonItem, err := findCommonChar(sackList[i], sackList[i-1], sackList[i-2])
			if err != nil {
				panic(err)
			}

			itemPrioritySum += getItemPriority(commonItem)
		}

	}

	return itemPrioritySum
}

func solve_day3() {
	filepath := "puzzle-inputs/day3-input"
	input := get_input(filepath)
	solution1 := solveDay3Part1(input)
	solution2 := solveDay3Part2(input)
	fmt.Println("Day 3: Rucksack Reorganization = ", solution1)
	fmt.Println("Day 3: Part Two =", solution2)
}
