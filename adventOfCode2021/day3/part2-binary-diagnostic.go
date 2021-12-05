package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode/utf8"
)

func main() {
	//wrangle the data
	var bin []string = input("data")

	//fuck with the slice until we're done
	var s []string = bin
	for b := 0; b < 12; b++ {

		//check if we're done
		if len(s) < 1 {
			fmt.Println("DONE: \n", s)
			return
		}

		//get the bitcriteria
		bitcriteria, _ := compute_shit(s)
		var debugString string = ""
		var a string = ""
		for i := len(s) - 2; i > 0; i-- {
			a = strings.Repeat("#", utf8.RuneCountInString(b-1))

			fmt.Println("bitcriteria: \t", a+bitcriteria[b])
			//remove entries from the slice that don't match the criteria
			if s[i][b] != bitcriteria[b] {
				fmt.Println("REMOVED: \t", s[i])
				s[i] = s[len(s)-1]
				s = s[:len(s)-1]
			} else {
				fmt.Println("APPROVED: \t", s[i])
			}

			//check if we're done
			if len(s) < 1 {
				fmt.Println("DONE: \n", s)
				return
			}
			a = ""

		}

	}
	fmt.Println(s)
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
		} else { //if even, treat the bit criteria as a "1"
			gamma += "0"
			epsilon += "1"
		}
		zeros = 0
		ones = 0
	}
	return gamma, epsilon
}
