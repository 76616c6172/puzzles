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
func data_wrangler(input string) ([]line, int) {
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
	return lines, line_amount
}

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

//Size of the SIZE by SIZE diagram
const SIZE = 1000

func main() {

	filename := "data"
	input := get_input(filename)

	//hold all the lines from the input
	lines, _ := data_wrangler(input)
	//hold all the points in the diagram
	points := make([][SIZE]point, SIZE)

	//Part2: Create a slice to store all the diagonal lines
	diag_lines := make([]line, 0)

	//remove all the lines that aren't horizontal or vertical
	for i := 0; i < len(lines); i++ {
		if lines[i].x1 != lines[i].x2 && lines[i].y1 != lines[i].y2 {
			diag_lines = append(diag_lines, lines[i]) //save the diag line
			lines[i] = lines[len(lines)-1]
			lines = lines[:len(lines)-1]
			i--
		}
	}
	//iterate through every line
	//update the board with every point of the line
	for i, _ := range lines {
		//draw vertical lines
		if lines[i].x1 == lines[i].x2 {
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

	//draw diagonal lines
	for _, v := range diag_lines {
		fmt.Println("start:", v.x1, v.y1)
		fmt.Println("end:", v.x2, v.y2)
		if v.y2 >= v.y1 {
			x := v.x1
			for y := v.y1; y <= v.y2; y++ {
				slope := float64((v.y2 - v.y1)) / float64((v.x2 - v.x1))
				points[y][x].count++
				x += int(slope)

			}
		} else {
			x := v.x1
			for y := v.y1; y >= v.y2; y-- {
				slope := float64((v.y2 - v.y1)) / float64((v.x2 - v.x1))
				fmt.Println(y, slope)
				fmt.Println("mark", x, y)
				points[y][x].count++
				x -= int(slope)

			}
		}

		fmt.Println()
	}

	//draw the board while calculating the answer
	var count int
	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			if points[y][x].count > 0 {
				fmt.Printf("%d", points[y][x].count)
				if points[y][x].count > 1 {
					count++
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("Part2 Answer: ", count)
	fmt.Println(diag_lines)
}
