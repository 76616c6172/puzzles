package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var BIT_AMOUNT = 5

func main() {
	//wrangle the data
	var bin []string = input_handler("example")
	//This slice will shrink
	var s []string = bin
	var oxygen_generator_rating int64
	//var	CO2_scrubber_rating int64

	//Compute oxygen_generator_rating
	bitcriteria, _ := compute_shit(s)
	for bit := 0; bit < BIT_AMOUNT; bit++ {
		for word := 0; word < len(s)-1; word++ {
			if s[word][bit] != bitcriteria[bit] {
				s[word] = s[len(s)-1]
				s = s[:len(s)-1]
				word = -1 //UGLY HACK: Since we just changed the order, start checking from index 0 again
			}
		}
		//then recompute the bitcriteria
		bitcriteria, _ = compute_shit(s)
		//UGLY HACK: For some reason the loop stops when there are only 2 entries left.
		if len(s) <= 2 {
			oxygen_generator_rating, _ = strconv.ParseInt(bitcriteria, 2, 32)
			break
		}
	}

	fmt.Println(oxygen_generator_rating)
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
