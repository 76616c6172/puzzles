package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	//wrangle the data
	var bin []string = input("data")

	//compute the data
	gamma, epsilon := compute_shit(bin)
	fmt.Println("Gamma:\n", gamma)
	fmt.Println("Epsilon:\n", epsilon)

	//convert binary strings to integers
	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 32)
	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 32)

	//compute and print the solution
	solution := gammaDecimal * epsilonDecimal
	fmt.Println("Solution\n", solution)
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
