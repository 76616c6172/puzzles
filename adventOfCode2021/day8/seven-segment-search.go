package main

import "fmt"

type segment struct {
	a bool
	b bool
	c bool
	d bool
	e bool
	f bool
	g bool
}

var numberedDisplays [10]segment

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
func render_horizontal_segments() {
	lineContent := " "
	for segment := 0; segment < 10; segment++ {
		if numberedDisplays[segment].a {
			lineContent += "####    "
		} else {
			lineContent += "....    "
		}
	}
	fmt.Println(lineContent)

}

// renders b,c,d,f segments
func render_vertical_segments() {
	lineContent := ""
	for segment := 0; segment < 10; segment++ {

		if numberedDisplays[segment].b {
			lineContent += "#    "
			if numberedDisplays[segment].c {
				lineContent += "#  "
			} else {
				lineContent += ".  "
			}
		} else {
			lineContent += ".    "
			if numberedDisplays[segment].c {
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
func render_display() {
	fmt.Println("0:      1:      2:      3:      4:      5:      6:      7:      8:      9:")
	render_horizontal_segments() // Render the a segments
	render_vertical_segments()   // render b and c segments
	render_horizontal_segments() // Render the d segments
	render_vertical_segments()   // render e and f segments
	render_horizontal_segments() // Render the g segments
}

// Solution for AoC2021 day8 problem: Seven segment search
func main() {

	// Turn all segments for testing
	//	for i, _ := range numberedDisplays {
	//		numberedDisplays[i].a = true
	//		numberedDisplays[i].b = true
	//		numberedDisplays[i].c = true
	//		numberedDisplays[i].d = true
	//		numberedDisplays[i].e = true
	//		numberedDisplays[i].f = true
	//		numberedDisplays[i].g = true
	//	}
	//
	render_display()

}
