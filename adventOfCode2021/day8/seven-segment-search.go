package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Each one of the 10 elements representss one seven segmented display
var numberedDisplay [10]map[string]bool

//type segment struct {
//	a bool
//	b bool
//	c bool
//	d bool
//	e bool
//	f bool
//	g bool
//}
//var numberedDisplay = map[string]bool{ "a": false, "b": false, "c": false, "d": false, "e": false, "f": false, "g": false}

// var employee = map[string]int{"Mark": 10, "Sandy": 20}

// 0:      1:      2:      3:      4:
// aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
// ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
// gggg    ....    gggg    gggg    ....
//
//  5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
// dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
// gggg    gggg    ....    gggg    gggg
//

// Prints the first line to stdout (a segments)
func render_horizontal_segments(display [10]map[string]bool, segment string) {
	lineContent := " "
	for displayNum := 0; displayNum < 10; displayNum++ {
		if display[displayNum][segment] {
			lineContent += "####    "
		} else {
			lineContent += "....    "
		}
	}
	fmt.Println(lineContent)

}

// renders b,c,d,f segments
func render_vertical_segments(display [10]map[string]bool, firstSegment string, secondSegment string) {
	lineContent := ""
	for displayNum := 0; displayNum < 10; displayNum++ {

		if display[displayNum][firstSegment] {
			lineContent += "#    "
			if display[displayNum][secondSegment] {
				lineContent += "#  "
			} else {
				lineContent += ".  "
			}
		} else {
			lineContent += ".    "
			if display[displayNum][secondSegment] {
				lineContent += "#  "
			} else {
				lineContent += ".  "

			}
		}
	}
	fmt.Println(lineContent)
	fmt.Println(lineContent)
}

// Renders out the seven segment display in stdout
func render_display(displays [10]map[string]bool) {
	fmt.Println("0:      1:      2:      3:      4:      5:      6:      7:      8:      9:")
	render_horizontal_segments(displays, "a")    // Render the a segments
	render_vertical_segments(displays, "b", "c") // render b and c segments
	render_horizontal_segments(displays, "d")    // Render the d segments
	render_vertical_segments(displays, "e", "f") // render e and f segments
	render_horizontal_segments(displays, "g")    // Render the g segments
}

// Builds and then returns a 2 dimensional array of the puzzle input
// Expects the input to the puzzle in the format:
//
//cdbga acbde eacdfbg adbgf gdebcf bcg decabf cg ebdgac egca | geac ceag faedcb cg
//edbcfag ebdca ebgad dagbef cfbed adcg dcgeab ac cae cgabef | gacd ac dgac ebdfc
//gceafb fcabedg ebfd bd ebacf bafcde daecbg dabfc abd acfgd | adb fagdc bd agfbec
//abd cebd bd gcbaf dcafb dface ecgfbda ecdfba fdegca dbafeg | dbecfag bd decb dbgfcea
//ecdagfb eb egfcb gcdbfa dcfeg aecb fabcg afebgd ebg agecbf | bge be abcfge bfedacg
func data_wrangler(filename string) [][2]string {
	var output [][2]string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var z [2]string
		line := scanner.Text()
		splitInput := strings.Split(line, " | ")
		z[0] = splitInput[0]
		z[1] = splitInput[1]

		output = append(output, z)
	}

	return output
}

// Resets the state of the seven segment displays back to off
func display_off(numberedDisplay [10]map[string]bool) {
	var letters = []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < 10; i++ { // Allocate memory to the maps
		for _, l := range letters {
			numberedDisplay[i][l] = false
		}
	}
}

// Takes as input the amount of segments turned on and prints out the possible displays
func XX(n int) {
	switch n {
	case 2:
		fmt.Println("number is 1 CF")
	case 3:
		fmt.Println("number is 7 ACF")
	case 4:
		fmt.Println("number is 4 BCDF")
	case 5:
		fmt.Println("number is 2 ACDEG, 3ACDFG, or 5ABDFG")
	case 6:
		fmt.Println("number is 0 ABCEFG, 6 BDEFG, or 9 ABCDFG")
	case 7:
		fmt.Println("number is 8 ABCDEFG")
	}
}

var puzzleInput [][2]string

// Solution for AoC2021 day8 problem: Seven segment search
func main() {

	// Create data structures
	for i := 0; i < 10; i++ { // Allocate memory to the maps
		numberedDisplay[i] = make(map[string]bool, 7)
	}

	render_display(numberedDisplay)

	// Read in the data
	puzzleInput = data_wrangler("example")

	// Testing output
	fmt.Println(puzzleInput[0][1])

	// Set all segments of all displays on for fun
	var letters = []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < 10; i++ { // Allocate memory to the maps
		for _, l := range letters {
			numberedDisplay[i][l] = true
		}
	}

}
