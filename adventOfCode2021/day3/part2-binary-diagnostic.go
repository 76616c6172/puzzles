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
	data, BIT_AMOUNT := input_handler("data")

	// make a copy of the data so we can still access the underlying data after we mangle it
	data2 := make([]string, len(data))
	copy(data2, data)

	// crunch numbers and completely mangle the array beneath the slice in the process
	oxygen_generator_rating := solve(data, BIT_AMOUNT, true)
	co2_scrubber_rating := solve(data2, BIT_AMOUNT, false)

	//Print the answer
	fmt.Println("Answer: \n", oxygen_generator_rating*co2_scrubber_rating)
}

// Computes the bicriteria for the current bit and returns true if the criteria is 1
func bit_is_1(arr []string, currentBit int) bool {
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
		return false
		//
	} else {
		return true
	}
}

func solve(arr []string, BIT_AMOUNT int, bit_rule bool) int64 {
	var criteria uint8
	for bit := 0; bit < BIT_AMOUNT; bit++ {
		fmt.Println(arr)
		//get bitcriteria
		if bit_is_1(arr, bit) == bit_rule {
			criteria = '1'
		} else {
			criteria = '0'
		}
		//safety check if we're done
		if len(arr) == 1 {
			answer, err := strconv.ParseInt(string(arr[0]), 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			return answer
		}
		for word := 0; word <= len(arr)-1; word++ {
			//if the criteria doesn't match, delete the word from the slice
			if arr[word][bit] != criteria {
				arr[word] = arr[len(arr)-1]
				arr = arr[:len(arr)-1]
				word-- //go back one element becaue we just shrunk the array
			}
		}
	}
	//if we were not done before, now we're definitely done!
	answer, err := strconv.ParseInt(string(arr[0]), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)
	return answer
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
