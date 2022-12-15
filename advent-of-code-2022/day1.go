package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getHighestCalories(s string) int {
	caloriesList := strings.Split(s, "\n")

	var highestCalorieNum int
	var currentElfCalories int

	for _, calories := range caloriesList {
		if calories != "" {
			calorieNum, err := strconv.Atoi(calories)
			if err != nil {
				panic(err)
			}
			currentElfCalories += calorieNum
		} else {
			if currentElfCalories > highestCalorieNum {
				highestCalorieNum = currentElfCalories
			}
			currentElfCalories = 0
		}
	}

	return highestCalorieNum
}

func addTop3Calories(s string) int {
	caloriesList := strings.Split(s, "\n")
	var currentElfCalories int
	var caloriesPerElf []int
	for _, calorie := range caloriesList {
		if calorie != "" {
			calorieNum, _ := strconv.Atoi(calorie)
			currentElfCalories += calorieNum
		} else {
			caloriesPerElf = append(caloriesPerElf, currentElfCalories)
			currentElfCalories = 0
		}
	}
	sort.Ints(caloriesPerElf)
	var sum int
	for i := 3; i > 0; i-- {
		sum += caloriesPerElf[len(caloriesPerElf)-i]
	}
	return sum
}

func solve_day1() {
	filepath := "puzzle-inputs/day1-input"
	input := get_input(filepath)
	fmt.Println("Day 1: Calorie Counting =", getHighestCalories(input))
	fmt.Println("Day 1: Part Two =", addTop3Calories(input))
}
