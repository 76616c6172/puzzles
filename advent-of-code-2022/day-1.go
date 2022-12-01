package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solve_day1() {
	filePath := "puzzle-inputs/day-1-part-1"
	input := getPuzzleInputFromFile(filePath)

	output := getMostCarriedCalories(input)
	fmt.Println("Day 1: Calorie Counting =", output)

	output = getSumOfTopThreeCalorieCarriers(input)
	fmt.Println("--- Part Two --- =", output)
	fmt.Println()
}

func getMostCarriedCalories(s string) int {
	listSeparatedByNewlines := strings.Split(s, "\n")

	var highestNumberOfCalories int
	var calsCurrentElf int

	for _, line := range listSeparatedByNewlines {
		if line != "" {
			numberOfCaloriesInCurrentLine, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			calsCurrentElf += numberOfCaloriesInCurrentLine
		} else {
			if calsCurrentElf > highestNumberOfCalories {
				highestNumberOfCalories = calsCurrentElf
			}
			calsCurrentElf = 0
		}
	}

	return highestNumberOfCalories
}

func getSumOfTopThreeCalorieCarriers(s string) int {
	listSeparatedByNewlines := strings.Split(s, "\n")

	var calsCurrentElf int
	var listOfCaloriesPerElf []int

	for _, line := range listSeparatedByNewlines {
		if line != "" {
			numberOfCaloriesInCurrentLine, _ := strconv.Atoi(line)
			calsCurrentElf += numberOfCaloriesInCurrentLine
		} else {
			listOfCaloriesPerElf = append(listOfCaloriesPerElf, calsCurrentElf)
			calsCurrentElf = 0
		}
	}

	sort.Ints(listOfCaloriesPerElf)
	var sum int
	for i := 3; i > 0; i-- {
		sum += listOfCaloriesPerElf[len(listOfCaloriesPerElf)-i]
	}
	return sum
}
