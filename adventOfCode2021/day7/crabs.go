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

func main() {
	filename := "input"

	crabs_arr := input_handler(filename)

	// ---
	// Idk exactly how to solve this
	// Let's implement a naive approach trying alignment on every possible slot while calculating the fuel cost of each

	// 1.) let's get the highest position a crab has
	highest := 0
	for _, v := range crabs_arr {
		if v > highest {
			highest = v
		}
	}
	fmt.Println("Input data:", crabs_arr)
	fmt.Println("Highest crab position:", highest)

	// 2.) Try to align on all possible positions and save the fuelcost

	// This map will hold the total fuelcost needed to align on every value in the format of
	// fueldata_m[value_we_tried_to_align_on]total_fuel_cost
	fueldata_m := make(map[int]int, 0)

	fuelcost := 0
	for i := 0; i < highest+1; i++ {
		fmt.Println("Trying to align on:", i)
		for crab := 0; crab < len(crabs_arr); crab++ {

			// Math happens here
			fuel_delta := math.Abs(float64(i - crabs_arr[crab]))
			fmt.Println(crabs_arr[crab], "costs", fuel_delta)
			fuelcost += int(fuel_delta)
		}
		fueldata_m[i] += fuelcost
		fmt.Println("Aligning on", i, "costs", fuelcost)
		fuelcost = 0 //reset for trying to align on next value
	}

	// Print the finished data from the brute force
	fmt.Println(fueldata_m)

	// 3.) Let's check which fuelcost is the smallest and print that as our answer
	lowest_cost := 4294967295
	for _, v := range fueldata_m {
		if v < lowest_cost {
			lowest_cost = v
		}
	}

	fmt.Println()
	fmt.Println("Lowest fuelcost was:", lowest_cost)
}
