package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	var gamma string
	var epsilon string
	var bin []string = input("data")

	ones := 0
	zeros := 0
	for a := 0; a < 12; a++ {
		for b := 0; b < len(bin)-1; b++ {
			//0
			if bin[b][a] == 48 {
				zeros++
			} else {
				ones++
			}
		}
		if zeros > ones {
			gamma += "0"
		} else if zeros < ones {
			gamma += "1"
		} else {
			log.Fatal("ERROR: Zeros and ones are probably exactly equal in this position.\n")
		}
		zeros = 0
		ones = 0
	}
	gammaSlice := strings.Split(gamma, "")
	for i := 0; i < 12; i++ {
		if gammaSlice[i] == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}

	// invert gamma to get epsilon
	// multiply
	// convert to int

	fmt.Println("Gamma:\n", gamma)
	fmt.Println("Epsilon:\n", epsilon)

	// FIXME: convert gamma and epsilon strings properly to REAL 12 bit binary data
	// multiply them and print the result
	fmt.Println("Now all that's left is to read both 12 bit chunks as decimal numbers and multiply them =O=")
}

func input(filename string) []string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	content := string(f)
	output := strings.Split(content, "\n")
	return output
}
