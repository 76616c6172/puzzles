package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var bin []string = input("data")
	gamma, epsilon := compute_shit(bin)

	//gammaSlice := strings.Split(gamma, "")
	//	for i := 0; i < 12; i++ {
	//		if gammaSlice[i] == "0" {
	//			epsilon += "1"
	//		} else {
	//			epsilon += "0"
	//		}
	//	}

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

func compute_shit(data []string) (string, string) {
	var gamma string = ""
	var epsilon string = ""
	ones := 0
	zeros := 0
	for a := 0; a < 12; a++ {
		for b := 0; b < len(data)-1; b++ {
			//0
			if data[b][a] == 48 {
				zeros++
			} else {
				ones++
			}
		}
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else if zeros < ones {
			gamma += "1"
			epsilon += "0"
		} else {
			log.Fatal("ERROR: Zeros and ones are probably exactly equal in this position.\n")
		}
		zeros = 0
		ones = 0
	}
	return gamma, epsilon
}
