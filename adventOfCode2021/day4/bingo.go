package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//TODO:
	//return winning numbers DONE.
	//return number of boards n
	//return 3 dimensional array boards[n][x][y]

	winning_nums := input_handler("example")
	fmt.Println(winning_nums)
}

func input_handler(filename string) []string {
	f, _ := ioutil.ReadFile(filename)
	a := string(f)

	//extract the winning numbers
	b := strings.Split(a, "\n")
	numbers := (strings.Split(b[0], ","))

	//extract the bingo boards somehow
	fmt.Println(b)

	return numbers
}
