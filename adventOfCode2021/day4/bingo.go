package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//TODO:
	//get number of boards n
	//read each board and store it in bingo_board_data
	//return winning numbers DONE.
	//create 3 dimensional slice bingo_board_data[n][x][y] DONE.

	//get the input from file
	winning_nums := input_handler("example")

	//bingo board data
	//bbd[board_number][y][x]
	BOARD_AMOUNT := 2
	bbd := make([][][]int, BOARD_AMOUNT)
	for a := range bbd {
		bbd[a] = make([][]int, 5)
		for b := range bbd[a] {
			bbd[a][b] = make([]int, 5)
		}

	}
	//test putting in this board
	//0 0 0 0 0
	//0 9 0 0 0
	//0 0 0 0 0
	//0 0 0 0 7
	//0 0 0 0 0
	bbd[0][1][1] = 9
	bbd[0][3][4] = 7

	//print out all the boards
	fmt.Println("Boards:")
	//print all boards
	for n := 0; n < BOARD_AMOUNT; n++ {
		for i := range bbd[n] {
			fmt.Println(bbd[n][i])
		}
		fmt.Println()
	}

	//print stuff
	fmt.Println()
	fmt.Println(winning_nums, "\n")
}

func input_handler(filename string) []string {
	f, _ := ioutil.ReadFile(filename)
	a := string(f)

	//extract the winning numbers
	b := strings.Split(a, "\n")
	numbers := (strings.Split(b[0], ","))

	//extract the bingo boards somehow and put them in the slice

	return numbers
}
