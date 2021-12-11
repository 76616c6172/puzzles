package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	//print out all the boards
	fmt.Println("Boards:")
	for n := 0; n < BOARD_AMOUNT; n++ {
		for i := range bbd[n] {
			fmt.Println(bbd[n][i])
		}
		fmt.Println()
	}

	fmt.Println("Numbers:", winning_nums)
	//number in y row
	//nuiy[current_board][row] = amount of matches in this row
	nuiy := make([][5]int, BOARD_AMOUNT) //matches in [board_index][y_row]
	nuix := make([][5]int, BOARD_AMOUNT) //matches in [board_index][x_row]

	for _, num := range winning_nums {

		//loop through the boards
		for cur_board := 0; cur_board < BOARD_AMOUNT; cur_board++ {

			//check for horizontal matchess
			for y := 0; y < 5; y++ {
				if row_contains(data[cur_board][y][:], num) {
					nuiy[cur_board][y]++
					fmt.Println("Number:", num, "found in board:", cur_board, "row:", y, "has", nuiy[cur_board][y], "matches!")
				}
				//check if any horizontal winners
				if nuiy[cur_board][y] == 5 {
					fmt.Println("Holy shit, Board", cur_board+1, "has 5 matches in row:", y)
					return
				}

				//check for vertical matches
				for x:= 0; x<5; x++ {
					if collum_contains(data[cur_board])

				}


			}

		}

	}

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
	log.Fatal("input_handler_2 error")
	return slice
}

func row_contains(row []string, e string) bool {
	for _, v := range row {
		if v == e {
			return true
		}
	}
	return false
}
