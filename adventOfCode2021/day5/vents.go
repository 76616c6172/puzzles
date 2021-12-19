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
	x1, y1 int //start point
	x2, y2 int //end point
}

// Stores data about each single point in the Diagram
type point struct {
	count int //how many times on line
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
				lines[i].x1 = x
				y, _ := strconv.Atoi(point[1])
				lines[i].y1 = y
			} else { //this is the end of the line segment
				x, _ := strconv.Atoi(point[0])
				lines[i].x2 = x
				y, _ := strconv.Atoi(point[1])
				lines[i].y2 = y

			}
		}
	}
	return lines
}

//Size of the SIZE by SIZE diagram
const SIZE = 1000

// return true if point is on line
func is_one_line(x int, y int, line line) bool {
	//do math to check if the point specified with x/y is on the supplied line
	//		x  y   				   	x1 y1  x2 y2
	//point: 2 1 is on line {[2 2] [2 1]}
	// TODO do math check
	if true {
		return true
	} else {
		return false
	}

}

func main() {

	filename := "data"
	input := get_input(filename)

	//hold all the lines from the input
	lines := data_wrangler(input)
	//hold all the points in the diagram
	points := make([][SIZE]point, SIZE)

	//remove all the lines that aren't horizontal or vertical
	for i := 0; i < len(lines); i++ {
		if lines[i].x1 != lines[i].x2 && lines[i].y1 != lines[i].y2 {
			lines[i] = lines[len(lines)-1]
			lines = lines[:len(lines)-1]
			i--
		}
	}
	//TODO: simpler solution:
	//iterate through every line
	//update the board with every point of the line
	for i, _ := range lines {
		//draw vertical lines
		if lines[i].x1 == lines[i].x2 {
			fmt.Println("VERTICAL LINE:", lines[i])
			y1 := lines[i].y1
			y2 := lines[i].y2
			for y := y2; y <= y1; y++ {
				points[y][lines[i].x2].count++
			}
			for y := y1; y <= y2; y++ {
				points[y][lines[i].x2].count++
			}
		}
		//draw horizontal lines
		if lines[i].y1 == lines[i].y2 {
			fmt.Println("HORIZONTAL LINE:", lines[i])
			x1 := lines[i].x1
			x2 := lines[i].x2
			for x := x2; x <= x1; x++ {
				points[lines[i].y2][x].count++
			}
			for x := x1; x <= x2; x++ {
				points[lines[i].y2][x].count++
			}
		}

	}
	fmt.Println(lines)

	/*
		// FIXME: Find a simpler solution than checking each point vs all lines
		// go through each point in the diagram and if it's one a line mark it with its counter
		for y := 0; y < SIZE; y++ {
			for x := 0; x < SIZE; x++ {
				for _, line := range lines {
					if is_one_line(x, y, line) {
						points[y][x].count++
					}
				}
			}
		}
	*/

	fmt.Println()

	//draw the board while calculating the answer
	var count int
	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			if points[y][x].count > 0 {
				fmt.Printf("%d", points[y][x].count)
				if points[y][x].count > 1 {
					count++
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}

	}
	fmt.Println("Part1 Answer: ", count)
}
