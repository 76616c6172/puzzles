package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Holds the data for each line segment
type line struct {
	start [2]int //line.start[x1][y1]
	end   [2]int //line.end[x2][y2]
}

// Reads file and returns the contents as a single string with the last newline removed
func get_input(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	result := strings.TrimSuffix(string(f), "\n")
	return string(result)
}

// returns a slice of all the lines
func data_wrangler(input string) []line {
	line_amount := strings.Count(input, "\n") + 1
	lines := make([]line, line_amount)
	one_line := strings.Split(input, "\n")

	//Store each point coordinate pair correctly in lines[] for each line
	for i := 0; i < line_amount; i++ {
		b := strings.Split(one_line[i], " -> ")
		for z, v := range b {
			point := strings.Split(v, ",")
			if z == 0 { //this is the start of the line segment
				x, _ := strconv.Atoi(point[0])
				lines[i].start[0] = x
				y, _ := strconv.Atoi(point[1])
				lines[i].start[1] = y
			} else { //this is the end of the line segment
				x, _ := strconv.Atoi(point[0])
				lines[i].end[0] = x
				y, _ := strconv.Atoi(point[1])
				lines[i].end[1] = y

			}
		}
	}
	return lines
}

func main() {

	filename := "data"
	input := get_input(filename)
	lines := data_wrangler(input)

	//print out all the stored lines
	for i, v := range lines {
		fmt.Println(i, v)
	}
}
