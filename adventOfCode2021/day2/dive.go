package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	//read in data from a textfile
	text := textInstructions("data")

	//separate the instructions into a slice
	instruction_slice := strings.Split(text, "\n")

	// loop through the instructions and compute the answer.
	// there are only 3 instructions:
	// forward +horizontal
	// up -depth
	// down +depth
	horizontal := 0
	depth := 0
	n := 0
	for _, b := range instruction_slice {
		if strings.Split(b, " ")[0] == "forward" {
			n, _ = (strconv.Atoi(strings.Split(b, " ")[1]))
			horizontal += n
		} else if strings.Split(b, " ")[0] == "up" {
			n, _ = (strconv.Atoi(strings.Split(b, " ")[1]))
			depth -= n
		} else if strings.Split(b, " ")[0] == "down" {
			n, _ = (strconv.Atoi(strings.Split(b, " ")[1]))
			depth += n
		}
	}

	solution := horizontal * depth

	fmt.Println(solution)
}

// read in a file and spit out instructions as a single long string
func textInstructions(filename string) string {

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	content := string(f)
	return content
}
