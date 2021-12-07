package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var BIT_AMOUNT = 5

func main() {
	//wrangle the data
	var bin []string = input_handler("example")
	//This slice will shrink
	var s []string = bin
	//	var o2 string
	fmt.Println(s)

	// FIXME: as the slice is shrinking we end up skipping some entries without checking if
	// they match the bitcriteria because they were scooted to a lower index in the slice
	bitcriteria_gamma, _ := compute_shit(s)
	for bit := 0; bit < BIT_AMOUNT; bit++ {
		for word := 0; word < len(s)-1; word++ {
			if s[word][bit] != bitcriteria_gamma[bit] {
				/* shrinking array without keeping the order

								s[word] = s[len(s)-1]
								s = s[:len(s)-1]
				:*/
				//shrinking while keeping order intact
				s = append(s[:word], s[word+1:]...)
			}
		}
		//then compute the bitcriteria
		bitcriteria_gamma, _ = compute_shit(s)
		fmt.Println(s)
	}
}

func input_handler(filename string) []string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	content := string(f)
	content = strings.TrimSuffix(content, "\n")
	output := strings.Split(content, "\n")
	return output
}

func compute_shit(data []string) (string, string) {
	var gamma string = ""
	var epsilon string = ""
	ones := 0
	zeros := 0
	for a := 0; a < BIT_AMOUNT; a++ {
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
		} else { //if even, treat the bit criteria as a "1"
			gamma += "1"
			epsilon += "0"
		}
		zeros = 0
		ones = 0
	}
	return gamma, epsilon
}
