package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	const DAYS int = 80

	fishies := get_input("data")
	fmt.Println(fishies)
	fmt.Println()

	for day := 0; day < DAYS; day++ {
		fmt.Println("day", day, fishies)
		for i := 0; i < len(fishies); i++ {
			fishies[i]--
			if fishies[i] < 0 {
				fishies[i] = 6
				fishies = append(fishies, 8+1)
			}

		}

	}

	fmt.Println("Part1 Answer:", len(fishies))
}
