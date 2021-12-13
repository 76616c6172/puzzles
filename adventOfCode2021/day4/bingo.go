package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//Returns winning_nums, board_board_3d, board_amount
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

func input_handler_2(slice [][][]string, raw_input []string, board_amount int) [][][]string {

	//extract the boards
	completed_boards := 0
	index := 2
	board_number := 0
	y := 0
	for completed_boards <= board_amount {
		row := strings.Split(raw_input[index], " ")
		for i, v := range row {
			if v == "" || v == " " || v == "\n" { //remove empty elements
				row = append(row[:i], row[i+1:]...)
			}
		}
		//DEBUG SANITY CHECK
		//fmt.Printf("Board board_3d: %#v \n", row) //instead of printing, store in board board_3d
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
	log.Fatal("input_handler_2 error")
	return slice
}

//returns true if the row contains the string e
func row_contains(row []string, e string) bool {
	for _, v := range row {
		if v == e {
			return true
		}
	}
	return false
}

// returns winner_index, winning_row_or_collum_index, last_called
func solve(board_3d [][][]string, winning_nums []string, board_amount int) (int, int, int, string, []string) {
	num_in_y := make([][5]int, board_amount) //matches in [board_index][y_row]
	num_in_x := make([][5]int, board_amount) //matches in [board_index][x_row]

	var last_called string
	var winning_y int = -1
	var winning_x int = -1
	var winner_index int

	new_winning_nums := winning_nums
	for i, num := range winning_nums {

		//loop through the boards
		for cur_board := 0; cur_board < board_amount; cur_board++ {

			//check for horizontal matchess
			for y := 0; y < 5; y++ {
				if row_contains(board_3d[cur_board][y][:], num) {
					num_in_y[cur_board][y]++
				}
				//check if any horizontal winners
				if num_in_y[cur_board][y] == 5 {
					fmt.Println("Holy shit, Board", cur_board+1, "has 5 matches in row:", y)
					winner_index = cur_board
					winning_y = y
					last_called = num
					new_winning_nums = winning_nums[:i+1]
					return winner_index, winning_x, winning_y, last_called, new_winning_nums
				}

			}
			//check for vertical matches
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if board_3d[cur_board][y][x] == num {
						num_in_x[cur_board][x]++
						if cur_board == 37 {
						}
					}
					//check if any horizontal winners
					if num_in_x[cur_board][x] == 5 {
						fmt.Println("Holy shit, Board", cur_board+1, "has 5 matches in collum:", x)
						winner_index = cur_board
						winning_x = x
						last_called = num
						new_winning_nums = winning_nums[:i+1]
						return winner_index, winning_x, winning_y, last_called, new_winning_nums
					}
				}
			}

		}

	}
	// returns winner_index, winning_row_or_collum_index, last_called
	return winner_index, winning_x, winning_y, last_called, new_winning_nums
}

//takes number and winning number, if number not part of winning number, return false
//FIXME: winning nums need to STOP once there is a winner
func marked(num string, winning_nums []string) bool {
	for _, v := range winning_nums {
		if v == num {
			fmt.Println("MATCH FOUND checking num:", num, "against win", v)
			return true
		}
	}
	return false
}
func main() {
	INPUT_FILE := "data"

	//get the input from file
	winning_nums, board_amount, raw_input := input_handler_1(INPUT_FILE)

	//create 3 dimensional slice to store all the boards
	//board_3d[board_index][y][x]
	board_3d := make([][][]string, board_amount)
	for a := range board_3d {
		board_3d[a] = make([][]string, 5)
		for b := range board_3d[a] {
			board_3d[a][b] = make([]string, 5)
		}

	}

	//store the boards
	board_3d = input_handler_2(board_3d, raw_input, board_amount)

	//Solve the puzzle by checking the boards for each number that is called and returning all this stuff:
	winner_index, winning_x, winning_y, last_called, new_winning_nums := solve(board_3d, winning_nums, board_amount)

	//print fhe winning board
	fmt.Println("Winning board:", winner_index+1)
	for i := range board_3d[winner_index] {
		fmt.Println(board_3d[winner_index][i])
	}

	// winner index is the right index
	//compute the answer
	fmt.Println(winning_nums)
	var result int
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !(marked(board_3d[winner_index][y][x], new_winning_nums)) {
				z, _ := strconv.Atoi(board_3d[winner_index][y][x])
				result += z
			}
		}
	}

	bla, _ := strconv.Atoi(last_called)
	ANSWER := result * bla
	fmt.Println("THE ANSWER :", ANSWER, "Last called: ", bla, "Board nums: ", result)
	fmt.Println(winning_y, winning_x)

}
