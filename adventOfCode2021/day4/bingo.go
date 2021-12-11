package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	//get the input from file
	winning_nums, board_data, BOARD_AMOUNT := input_handler("example")

	//bingo board data
	//bbd[board_index][y][x]
	bbd := make([][][]int, BOARD_AMOUNT)
	for a := range bbd {
		bbd[a] = make([][]int, 5)
		for b := range bbd[a] {
			bbd[a][b] = make([]int, 5)
		}

	}

	/*
		for i, v := range board_data {
			fmt.Println(i, v)
		}
	*/

	//Wonderful we have a slice of rows but the fucking whitespace man
	board_index := 0
	inside := true
	y := 0
	count := 0
	//loop through rows
	for _, row := range board_data {
		x := strings.Split(row, " ")
		for _, v := range x {
			fmt.Println("Element is", v, "Count: ", count)
			if y > 4 {
				inside = false
			}
			if count < 5 && v != " " && inside {
				number, _ := strconv.Atoi(v)
				bbd[board_index][y][count] = number
				count++
				fmt.Println("ADDED: ", number, "Increased count to: ", count)
			}
			if count > 4 && v == "" && y < 5 {
				y++
				count = 0
			}
		}
	}
	//PRINTING
	//print out all the boards
	fmt.Println("Boards:")
	for n := 0; n < BOARD_AMOUNT; n++ {
		for i := range bbd[n] {
			fmt.Println(bbd[n][i])
		}
		fmt.Println()
	}

	//we can loop over the winning numbers like this:
	for _, x := range winning_nums {
		fmt.Println(x)
	}
}

//Returns winning_nums, board_data, board_amount
func input_handler(filename string) ([]string, []string, int) {
	f, _ := ioutil.ReadFile(filename)
	a := string(f)

	//extract the amount of boards
	newLines := strings.Count(a, "\n")
	board_amount := (newLines - 1) / 6

	//extract the winning numbers
	b := strings.Split(a, "\n")
	numbers := (strings.Split(b[0], ","))

	//
	board_data := b[2:]
	//extract the bingo boards somehow and put them in the slice
	return numbers, board_data, board_amount
}
