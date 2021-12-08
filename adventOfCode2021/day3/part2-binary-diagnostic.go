package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//wrangle the data
	data, BIT_AMOUNT := input_handler("example")

	//make a copies of the original array, so we can shrink it later
	ox := data
	fmt.Println(ox)
	for bit := 0; bit < BIT_AMOUNT; bit++ {
		criteria := bitcriteria(ox, bit)
		//for i, v := range ox { //INTERESTING: when looping over the range reducing the index did not push back loop
		for word := 0; word < len(ox)-1; word++ {
			if ox[word][bit] != criteria {
				ox[word] = ox[len(ox)-1] //THEORY: Is the problem that we're trying to replace the last index?
				ox = ox[:len(ox)-1]
				word-- //go back one element becaue we just shrunk the array
			}
			fmt.Println(ox)
		}
	}

	//figure out the oxygen_generator_rating
	//figure out the co2_scrubber_rating
	//print the slice
}

// Computes the bicriteria for the current bit
func bitcriteria(arr []string, currentBit int) uint8 {
	var zeros int
	var ones int
	for _, v := range arr {
		if v[currentBit] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		return '0'
	} else {
		return '1'
	}
}

// Takes in the data and returns a slice and the amount of bits
func input_handler(filename string) ([]string, int) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(f)
	content = strings.TrimSuffix(content, "\n")
	output := strings.Split(content, "\n")
	bitsize := len(output[0])
	return output, bitsize
}
