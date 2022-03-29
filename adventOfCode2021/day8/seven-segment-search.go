package main

import "fmt"

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

// Solution for AoC2021 day8 problem: Seven segment search
func main() {

	// Create data structures
	var numberedDisplay [10]map[string]bool
	for i := 0; i < 10; i++ { // Allocate memory to the maps
		numberedDisplay[i] = make(map[string]bool, 7)
	}

	// Set all segments of all displays on for fun
	var letters = []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < 10; i++ { // Allocate memory to the maps
		for _, l := range letters {
			numberedDisplay[i][l] = true
		}
	}

	render_display(numberedDisplay)

}
