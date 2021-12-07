package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//wrangle the data
	var bin []string = input_handler("testdata")
	fmt.Printf("Og Slice \t: %#v\n", bin)

	//fuck with the slice until we're done
	var s []string = bin
	fmt.Printf("Cp slice\t: %#v\n", s)
	//	var o2 string

	for bit := 0; bit < 12; bit++ {
		bitcriteria_gamma, _ := compute_shit(s)
		fmt.Println("------ CHECKING BIT ------", bit)
		fmt.Println("criteria: \t", bitcriteria_gamma)

		for a := 0; a < len(s)-1; a++ {
			if s[a][bit] != bitcriteria_gamma[bit] {
				fmt.Println("rejected: \t", s[a])
				s = append(s[:a], s[a+1:]...)
				fmt.Println("SLICE LENGTH: ", len(s))
			}
			/*
				o2 = s[a]
				o2Decimal, _ := strconv.ParseInt(o2, 2, 32)
				fmt.Println("o2Decimal:", o2Decimal)
			*/
		}
		fmt.Printf("CURRENT ARRAY%#v\n", s)
	}
	fmt.Println(s)

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
			gamma += "1"
			epsilon += "0"
		}
		zeros = 0
		ones = 0
	}
	return gamma, epsilon
}
