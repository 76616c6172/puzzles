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

func solve(fishies []int, days int) int {
	// Compute solution
	for day := 0; day < days; day++ {
		//fmt.Println("day", day, fishies)
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

	fishies := get_input("example")
	fishies_part2 := make([]int, len(fishies))
	copy(fishies_part2, fishies)

	part1_answer := solve(fishies, DAYS_PART_ONE)
	fmt.Println("Part1 Answer:", part1_answer)

	// FIXME: Running out of memory, reduce memory consumption!?
	part2_answer := solve(fishies_part2, DAYS_PART_TWO)
	fmt.Println("Part2 Answer:", part2_answer)

}
