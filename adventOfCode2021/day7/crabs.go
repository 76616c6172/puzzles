package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// take in input data and return as slice
func input_handler(filename string) []int {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// I love parsing strings..
	str1 := strings.TrimSuffix(string(f), "\n")
	sl1 := strings.Split(str1, ",")

	crabs_arr := make([]int, 0)

	// put the position of each crab into the slice
	for _, v := range sl1 {
		num, _ := strconv.Atoi(v)
		crabs_arr = append(crabs_arr, num)
	}

	return crabs_arr
}

func solve_part1(crabs_arr []int) int {
	// ---
	// Idk exactly how to solve_part1 this
	// Let's implement a naive approach trying alignment on every possible slot while calculating the fuel cost of each

	// 1.) let's get the highest position a crab has
	highest := 0
	for _, v := range crabs_arr {
		if v > highest {
			highest = v
		}
	}
	//fmt.Println("Input data:", crabs_arr)
	//fmt.Println("Highest crab position:", highest)

	// 2.) Try to align on all possible positions and save the fuelcost

	// This map will hold the total fuelcost needed to align on every value in the format of
	// fueldata_m[value_we_tried_to_align_on]total_fuel_cost
	fueldata_m := make(map[int]int, 0)

	fuelcost := 0
	for i := 0; i < highest+1; i++ {
		//fmt.Println("Trying to align on:", i)
		for crab := 0; crab < len(crabs_arr); crab++ {

			// Math happens here
			fuel_delta := math.Abs(float64(i - crabs_arr[crab]))
			//fmt.Println(crabs_arr[crab], "costs", fuel_delta)
			fuelcost += int(fuel_delta)
		}
		fueldata_m[i] += fuelcost
		//fmt.Println("Aligning on", i, "costs", fuelcost)
		fuelcost = 0 //reset for trying to align on next value
	}

	// Print the finished data from the brute force
	//fmt.Println(fueldata_m)

	// 3.) Let's check which fuelcost is the smallest and print that as our answer
	lowest_cost := 4294967295
	for _, v := range fueldata_m {
		if v < lowest_cost {
			lowest_cost = v
		}
	}
	return lowest_cost
}

func solve_part2(crabs_arr []int) int {

	// 1.) let's get the highest position a crab has
	highest := 0
	for _, v := range crabs_arr {
		if v > highest {
			highest = v
		}
	}
	//fmt.Println("Input data:", crabs_arr)
	//fmt.Println("Highest crab position:", highest)

	// 2.) Try to align on all possible positions and save the fuelcost

	// This map will hold the total fuelcost needed to align on every value in the format of
	// fueldata_m[value_we_tried_to_align_on]total_fuel_cost
	fueldata2_m := make(map[int]int, 0)

	//fuelcost := 0
	movement_steps := 0 //for part 2
	cost_per_step := 1

	for i := 0; i < highest+1; i++ {
		//i := 5
		//fmt.Println("Aligning on:", i)
		for crab := 0; crab < len(crabs_arr); crab++ {
			//fmt.Println(crabs_arr)

			// Math happens here
			steps_needed := math.Abs(float64(i - crabs_arr[crab]))
			//fmt.Println("Crab", crabs_arr[crab], "needs :", steps_needed, "steps")

			//fmt.Println("TOTAL fuel cost before adding the crab", crab, "is", fueldata2_m[i])
			movement_steps += int(steps_needed)
			for x := 0; x < int(steps_needed); x++ {
				//fmt.Println("Adding cost of", cost_per_step)
				fueldata2_m[i] += cost_per_step
				cost_per_step++
			}
			//fmt.Println("Fuel cost after adding the crab", crab, "is", fueldata2_m[i])

			//reset both before we go on to the next crab!
			//fuelcost = 0
			cost_per_step = 1
		}
	}

	// Print the finished data from the brute force
	//fmt.Println(fueldata_m)

	// 3.) Let's check which fuelcost is the smallest and print that as our answer
	lowest_cost := 4294967295
	for _, v := range fueldata2_m {
		if v < lowest_cost {
			lowest_cost = v
		}
	}
	return lowest_cost
}

func main() {
	filename := "input"

	// Wrangle data
	crabs_arr := input_handler(filename)

	// Solve part1
	p1_answer := solve_part1(crabs_arr)
	fmt.Println("Part1, lowest fuelcost is:", p1_answer)

	// Solve part2 - FIXME: it works but the approach is very naive and is therefore very slow :c
	p2_answer := solve_part2(crabs_arr)
	fmt.Println("Part2, lowest extended fuelcost is:", p2_answer)
}
