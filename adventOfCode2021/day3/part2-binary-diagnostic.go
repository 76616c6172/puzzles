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
	data, BIT_AMOUNT := input_handler("example")

	//copy the data
	data1 := data
	data2 := data
	//this is my code
	//so here we are just printing out the data which is a dynamic array
	// HERE:
	fmt.Println("d1\t", data1)
	fmt.Println("d2\t", data2)

	//make a copies of the original array, so we can shrink it later
	oxygen_generator_rating := find_oxygen_generator_rating(data1, BIT_AMOUNT)
	co2_scrubber_rating := find_co2_scrubber_rating(data2, BIT_AMOUNT)
	excuse_me(data)
	excuse_you(data)

	// SO ANYWAYS CHECK THIS OUT BRANDON
	// I WILL JUST MOVE THE FUNCTION CALL UP

	//Print the answer
	fmt.Println(oxygen_generator_rating * co2_scrubber_rating)

}

// all excuse_me and excuse_you functions do is, take in the slice and print it
// so it should be the same as what I am giving it?
func excuse_me(x []string) {
	fmt.Println("ex \t", x)
}
func excuse_you(x []string) {
	fmt.Println("ex u \t", x)
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

// Computes the bicriteria for the current bit
func reverse_bitcriteria(arr []string, currentBit int) uint8 {
	var zeros int
	var ones int
	for _, v := range arr {
		if v[currentBit] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	if ones < zeros {
		return '1'
	} else {
		return '0'
	}
}

func find_oxygen_generator_rating(ox []string, BIT_AMOUNT int) int64 {
	//fmt.Println("oxy\t", ox)
	// I am not touching the input
	for bit := 0; bit < BIT_AMOUNT; bit++ {
		criteria := bitcriteria(ox, bit)
		//for i, v := range ox { //INTERESTING: when looping over the range reducing the index did not push back loop
		for word := 0; word < len(ox)-1; word++ {
			if len(ox) == 2 { //check if we're done
				var answer []byte
				for g := 0; g < BIT_AMOUNT; g++ {
					answer = append(answer, bitcriteria(ox, g))
					//fmt.Printf("%c", bitcriteria(ox, g))
				}
				//oxygen_generator_rating_binary, _ := strconv.Atoi(string(answer))
				oxygen_generator_rating, _ := strconv.ParseInt(string(answer), 2, BIT_AMOUNT+1) //Q: Why BIT_AMOUNT+1 needed?
				return oxygen_generator_rating
			}
			if ox[word][bit] != criteria {
				ox[word] = ox[len(ox)-1] //THEORY: Is the problem that we're trying to replace the last index?
				ox = ox[:len(ox)-1]
				word-- //go back one element becaue we just shrunk the array
			}
		}
	}
	log.Fatal("oxygen_generator_rating not found")
	return -1
}

func find_co2_scrubber_rating(co []string, BIT_AMOUNT int) int64 {
	//fmt.Println("co2\t", co)
	for bit := 0; bit < BIT_AMOUNT; bit++ {
		//fmt.Println(co)
		criteria := reverse_bitcriteria(co, bit)
		//fmt.Println("\tchecking", bit, "index")
		//fmt.Println("\tCriteria is:", string(criteria))
		//for i, v := range ox { //INTERESTING: when looping over the range reducing the index did not push back loop
		for word := 0; word < len(co)-1; word++ {
			if len(co) == 2 { //check if we're done
				var answer []byte
				for g := 0; g < BIT_AMOUNT; g++ {
					//THEORY: reverse bitcriteria is fucked?
					answer = append(answer, reverse_bitcriteria(co, g))
					//fmt.Printf("%c", bitcriteria(co, g))
				}
				//co, _ := strconv.Atoi(string(answer))
				co2_generator_rating, _ := strconv.ParseInt(string(answer), 2, BIT_AMOUNT+1) //Q: Why BIT_AMOUNT+1 needed?
				return co2_generator_rating
			}
			if co[word][bit] != criteria {
				//fmt.Println("REMOVING:", co[word])
				co[word] = co[len(co)-1] //THEORY: Is the problem that we're trying to replace the last index?
				co = co[:len(co)-1]
				word-- //go back one element becaue we just shrunk the array
			}
		}
	}
	log.Fatal("co2_scrubber_rating_not_found")
	return -1
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
