package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Opens file and returns string
func get_input(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(f)
}

// Holds information about a particular board, such as its numbers
type board struct {
	//board.nums[y][x]
	num                 [5][5]int  //stores the number in the specified cell
	crossed             [5][5]bool //stores if a cell was crossed off
	flag                bool       //flag is true if the board already won
	crossed_in_y_row    [5]int
	crossed_in_x_collum [5]int
	ranking             int //stores the order in which the board won
	last_called_with    int //stores the number that was called that made the board win
}

// Returns winning numbers and all board data
func data_wrangler(input string) ([]int, []board, int) {
	s := strings.Split(input, "\n")
	nums := strings.Split(s[0], ",")
	amount_winning_nums := len(nums)

	//extract the winning numbers
	winning_nums := make([]int, amount_winning_nums)
	for i, v := range nums {
		a := string(v)
		b, _ := strconv.Atoi(a)
		winning_nums[i] = b
	}

	//extract the boards
	amount_of_boards := (strings.Count(input, "\n") - 1) / 6
	boards := make([]board, amount_of_boards)
	x, y, n := 0, 0, 0
	for _, row := range s[1:] {
		r := strings.Split(row, " ")
		for _, e := range r {
			if e != "" {
				num, err := strconv.Atoi(e)
				if err != nil {
					log.Fatal(err)
				}
				boards[n].num[y][x] = num
				x++
				if x == 5 {
					y++
					x = 0
				}
				if y == 5 {
					n++
					y = 0
				}
			}
		}
	}

	return winning_nums, boards, amount_of_boards
}

func main() {
	FILENAME := "data"
	input := get_input(FILENAME)
	winning_nums, boards, amount_of_boards := data_wrangler(input)

	//print winning numbers
	fmt.Println(winning_nums)

	//iterate through the numbers and mark off the boards
	rank := 1 //the first board that wins will be rank 1, the next rank2 and so on
	for _, winning_num := range winning_nums {
		//iterate through the boards
		for board_index := 0; board_index < amount_of_boards; board_index++ {
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if !boards[board_index].flag && winning_num == boards[board_index].num[y][x] {
						boards[board_index].crossed[y][x] = true
						boards[board_index].crossed_in_x_collum[x]++
						boards[board_index].crossed_in_y_row[y]++
						boards[board_index].last_called_with = winning_num

						//check for winners horizontally
						if boards[board_index].crossed_in_x_collum[x] == 5 {
							boards[board_index].ranking = rank
							boards[board_index].flag = true

							rank++
						}
						//check for winners vertically
						if boards[board_index].crossed_in_y_row[y] == 5 {
							boards[board_index].ranking = rank
							boards[board_index].flag = true
							rank++
						}
					}
				}
			}
		}
	}
	//print out the boards nicely:
	for n := 0; n < amount_of_boards; n++ {
		fmt.Printf("Board rank: %d\n", boards[n].ranking)
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if boards[n].crossed[y][x] {
					fmt.Printf("⦗%d⦘\t", boards[n].num[y][x])
				} else {
					fmt.Printf("%d\t", boards[n].num[y][x])
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	//Part1:
	//compute board score of winner
	proof := 0
	for n := 0; n < amount_of_boards; n++ {
		if boards[n].ranking == 1 {
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if !boards[n].crossed[y][x] {
						proof += boards[n].num[y][x]
					}
				}
			}
			proof = proof * boards[n].last_called_with
		}
	}
	fmt.Println("(Part1) Proof of winning board is:", proof)

	//Part2:
	//find the board that wins last

	r := 0           //rank placeholder
	last_index := -1 //index of last board
	for n := 0; n < amount_of_boards; n++ {
		if boards[n].ranking > r {
			last_index = n
			r = boards[n].ranking
		}
	}

	//compute the proof for the last board
	proof = 0 //reset the number from part1
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !boards[last_index].crossed[y][x] {
				proof += boards[last_index].num[y][x]
			}
		}
	}
	proof = proof * boards[last_index].last_called_with

	fmt.Println("(Part2) Proof of Last board is:", proof)
}
