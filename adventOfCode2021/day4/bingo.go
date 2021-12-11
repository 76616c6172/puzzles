package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	//get the input from file
	winning_nums, BOARD_AMOUNT, raw := input_handler_1("example")

	//bingo board data
	//bbd[board_index][y][x]
	bbd := make([][][]string, BOARD_AMOUNT)
	for a := range bbd {
		bbd[a] = make([][]string, 5)
		for b := range bbd[a] {
			bbd[a][b] = make([]string, 5)
		}

	}

	data := input_handler_2(bbd, raw, BOARD_AMOUNT)

	/*
		for i, v := range board_data {
			fmt.Println(i, v)
		}
	*/

	//Wonderful we have a slice of rows but the fucking whitespace man

	/*
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
	*/
	//PRINTING
	fmt.Println(data, "\n\n")

	//print out all the boards
	fmt.Println("Boards:")
	for n := 0; n < BOARD_AMOUNT; n++ {
		for i := range bbd[n] {
			fmt.Println(bbd[n][i])
		}
		fmt.Println()
	}

	/*
		//we can loop over the winning numbers like this:
		for _, x := range winning_nums {
			fmt.Println(x)
		}
	*/
	fmt.Println("Numbers:", winning_nums)
}

//Returns winning_nums, board_data, board_amount
func input_handler_1(filename string) ([]string, int, []string) {
	f, _ := ioutil.ReadFile(filename)
	a := string(f)

	//extract the amount of boards
	newLines := strings.Count(a, "\n")
	board_amount := (newLines - 1) / 6

	//extract the winning numbers
	b := strings.Split(a, "\n")
	numbers := (strings.Split(b[0], ","))

	//extract the bingo boards somehow and put them in the slice
	return numbers, board_amount, b
}
func input_handler_2(slice [][][]string, raw []string, board_amount int) [][][]string {

	//extract the boards
	completed_boards := 0
	index := 2
	board_number := 0
	y := 0
	for completed_boards <= board_amount {
		row := strings.Split(raw[index], " ")
		for i, v := range row {
			if v == "" || v == " " || v == "\n" { //remove empty elements
				row = append(row[:i], row[i+1:]...)
			}
		}
		//DEBUG SANITY CHECK
		//fmt.Printf("Board data: %#v \n", row) //instead of printing, store in board data
		//fmt.Println("BOARD:", board_number)
		for x, v_2 := range row { //put the current row into the board
			slice[board_number][y][x] = v_2
		}
		y++
		index++
		if y == 6 {
			y = 0
			board_number++
		}
		if board_number > board_amount {
			return slice
		}
		if index > board_amount*6 {
			return slice
		}
	}
	//PLACEHOLDER RETURN
	return slice
}
