package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const DISPLAY_AMOUNT int = 10

var numberedDisplay [10]map[string]bool // Each one of the 10 elements representss one seven segmented display
var puzzleInput [][2]string             // Holds the data of the puzzle
var easyDigitCounter int = 0            // Counts the easy digits from the input data
var Encoding = map[string]int{"acedgfb": 8,
	"cdfbe":  5,
	"gcdfa":  2,
	"fbcad":  3,
	"dab":    7,
	"cefabd": 9,
	"cdfgeb": 6,
	"eafb":   4,
	"cagedb": 0,
	"ab":     1,
}

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
// cdbga acbde eacdfbg adbgf gdebcf bcg decabf cg ebdgac egca | geac ceag faedcb cg
// edbcfag ebdca ebgad dagbef cfbed adcg dcgeab ac cae cgabef | gacd ac dgac ebdfc
// gceafb fcabedg ebfd bd ebacf bafcde daecbg dabfc abd acfgd | adb fagdc bd agfbec
// abd cebd bd gcbaf dcafb dface ecgfbda ecdfba fdegca dbafeg | dbecfag bd decb dbgfcea
// ecdagfb eb egfcb gcdbfa dcfeg aecb fabcg afebgd ebg agecbf | bge be abcfge bfedacg
func get_input_from_data(filename string) [][2]string {
	var output [][2]string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var splitLine [2]string
		line := scanner.Text()
		splitInput := strings.Split(line, " | ")
		splitLine[0] = splitInput[0]
		splitLine[1] = splitInput[1]

		output = append(output, splitLine)
	}

	return output
}

// Set all segments of all displays on for fun
func display_on() {

	var letters = []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < 10; i++ { // Allocate memory to the maps
		for _, l := range letters {
			numberedDisplay[i][l] = true
		}
	}

}

// Resets the state of the seven segment displays back to off
func display_off() {
	var letters = []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < 10; i++ { // Allocate memory to the maps
		for _, l := range letters {
			numberedDisplay[i][l] = false
		}
	}
}

// Checks for easy strings that can be identified purely based on length (2, 3, 4, or 7)
// As a side effect it adds +1 to the global easyDigitCounter everytime it identifies an easy digit
// Returns true or false and the identified digit (returns -1 if digit can't be identified)
func get_num_from_ez_length(length int) (bool, string) {
	switch length {
	case 2:
		easyDigitCounter++
		return true, "1" //number is 1
	case 3:
		easyDigitCounter++
		return true, "7" //number is 7
	case 4:
		easyDigitCounter++
		return true, "4" //number is 4
	case 5:
		return false, "235" // number is either 2 , 3, 5
	case 6:
		return false, "069" // number is either 0 , 6, 9
	case 7:
		easyDigitCounter++
		return true, "8" //number is 8
	}
	return false, ""
}

// Returns the correct number from a hard pattern and true
// Returns false and 99 if no number could be found
func getDigitFromHardPattern(input string) (string, bool) {
	// Example:
	// 						cefdb can be 2, 3 or 5
	// sorted:		bcdef
	var key = map[string]string{
		"abcdeg": "-1", // this is actually 0
		"acdfg":  "2",
		"abcdf":  "3",
		"bcdef":  "5",
		"bcdefg": "6",
		"abcde":  "8",
		"abcdef": "9",
	}

	// Encode the real 0 as -1 so we can assume 0 means not found
	if key[input] == "" {
		return "_", false
	}

	return key[input], true
}

// Solves part one of the puzzle and prints the solution
func solve_part_one() {
	for _, line := range puzzleInput {
		//signalPattern := strings.Fields(line[0])
		outputValue := strings.Fields(line[1])

		// Parse each wrong display pattern in the output value
		for _, str := range outputValue {
			fmt.Println("Parsing: ", str)
			isEasy, _ := get_num_from_ez_length(len(str))
			fmt.Println(isEasy)
		}
	}

	fmt.Println("Answer to part one is:", easyDigitCounter, "\n")
}

// Returns an alphabetically sorted version of the input string
func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

// Needed for sortString
type sortRuneString []rune

// Needed for sortString
func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Needed for sortString
func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

// Needed for sortString
func (s sortRuneString) Len() int {
	return len(s)
}

var partTwoAnswer int

// Solves part two of the puzzle and prints the solution to stdout
func solve_part_two() {

	// iterate over every line from the input
	for _, line := range puzzleInput {

		var numberSlice []string // stores the decoded number for each output value

		signal := strings.Fields(line[1]) // iterate only over the fields in the second half of the data ("")
		for _, entry := range signal {

			isEasyPattern, numberStr := get_num_from_ez_length(len(entry)) // Check for easy pattern

			if isEasyPattern {
				numberSlice = append(numberSlice, numberStr)
				fmt.Println("EZ:", numberStr)

			} else {

				sortedEntry := sortString(entry) // Sort entry and match against map for hard patterns
				numberStr, match := getDigitFromHardPattern(sortedEntry)

				if match {
					numberSlice = append(numberSlice, numberStr)
					fmt.Println("Hard number:", numberStr)
					continue

				}
				numberSlice = append(numberSlice, "_")
				fmt.Println("Can't find number:", entry)
			}

		}
		// Get the end result
		solutionStr := strings.Join(numberSlice, "")
		fmt.Println(numberSlice)
		if strings.Contains(solutionStr, "_") {
			fmt.Println("Skipping string with missing numbers")
		} else {
			number, err := strconv.Atoi(solutionStr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(number)
			partTwoAnswer += number
		}

	}
	fmt.Println("partTwoAnswer top part 2:", partTwoAnswer)
}

func main() {

	for i := 0; i < DISPLAY_AMOUNT; i++ { // Allocate memory to the maps
		numberedDisplay[i] = make(map[string]bool, 7)
	}

	puzzleInput = get_input_from_data("example")

	solve_part_one()
	solve_part_two()

	/*
		display_on()                    // Turn on just for fun
		render_display(numberedDisplay) // Print displays just for fun
		display_off()
		render_display(numberedDisplay) // Print displays just for fun
	*/
}
