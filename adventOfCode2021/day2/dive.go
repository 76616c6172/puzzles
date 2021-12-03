package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	//read in data from a textfile
	text := textInstructions("data")

	instructions := make(map[string]int)

	fmt.Println(text)
	fmt.Println(instructions)

	// take in the data
	//fmt.Printf("%T", instructions)

	/*
		// compute depth and horizontal values
		horizontal := 0
		depth := 0

		// there are only 3 instructions:
		// forward +horizontal
		// up -depth
		// down +depth
		// no backwards!




		solution := horizontal  * depth
		fmt.Println(solution)
	*/
}

// read in a file and spit out instructions as a single long string
func textInstructions(filename string) string {

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	content := string(f)

	/*
		d := strings.Split(content, "\n")

		var result []int
		var number int
		for _, s := range d {
			number, _ = strconv.Atoi(s)
			result = append(result, number)
		}
	*/

	return content
}
