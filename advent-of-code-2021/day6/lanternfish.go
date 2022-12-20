package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Spits out a slice with the fishies
func get_input(filename string) []int {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	//hold the data of the fish
	arr := make([]int, 0)

	f_str := strings.TrimSuffix(string(f), "\n")
	input := strings.Split(f_str, ",")
	for _, e := range input {
		num, _ := strconv.Atoi(e)
		arr = append(arr, num)
	}
	return arr
}

// Computes solution
func solve(fishies []int, days int) int {
	for day := 0; day < days; day++ {
		for i := 0; i < len(fishies); i++ {
			fishies[i]--
			if fishies[i] < 0 {
				fishies[i] = 6
				fishies = append(fishies, 8+1)
			}
		}
		if day%32 == 0 {
			fmt.Println("At day", day, "Fish:", len(fishies))
		}
	}
	return len(fishies)
}

func main() {
	const DAYS_PART_ONE int = 80
	const DAYS_PART_TWO int = 256
	const INPUT_FILENAME string = "data"

	fishies := get_input(INPUT_FILENAME)
	fishies_part2 := make([]int, len(fishies))
	copy(fishies_part2, fishies)

	part1_answer := solve(fishies, DAYS_PART_ONE)
	fmt.Println("Part1 Answer:", part1_answer)

	// PART TWO:
	/*
		part2_answer := solve(fishies_part2, DAYS_PART_TWO)
		fmt.Println("Part2 Answer:", part2_answer)
	*/

	fish_map := make(map[int]int)
	for _, v := range fishies_part2 {
		fish_map[v]++
	}

	for day := 0; day < DAYS_PART_TWO; day++ {
		fish_map[9] = fish_map[0]  //neccessary placeholder
		fish_map[7] += fish_map[0] //the newly hatched fishies go on a 7 day timer that will become 6 after the loop
		for i := 0; i < 9; i++ {
			fish_map[i] = fish_map[i+1]
		}
		fish_map[9] = 0
		//fmt.Println("day:", day+1, fish_map)
	}

	part2_answer := 0
	for i := 0; i < 9; i++ {
		part2_answer += fish_map[i]
	}
	fmt.Println("Part2: Answer:", part2_answer)

}
