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
	var bin []string = input_handler("data")

	//fuck with the slice until we're done
	var s []string = bin
	var o2 string

	for bit := 0; bit < 12; bit++ {
		//get the bitcriteria
		bitcriteria_gamma, _ := compute_shit(s)
		fmt.Println(bit, "<-- INDEX THE BIT WE'RE CHECKING FOR CRITERIA")
		fmt.Println("Bitcriteria g: \t", bitcriteria_gamma)
		for a := 0; a < len(s)-1; a++ {
			fmt.Println("Chunk we're on: \t", s[a])

			if s[a][bit] != bitcriteria_gamma[bit] {
				fmt.Println("------- SHRINKING THE ARRAY ---------")
				s = append(s[:a], s[a+1:]...)
			}

			o2 = s[a]
			if bit == 12 && a < 2 {
				fmt.Println(s)
			}
			o2Decimal, _ := strconv.ParseInt(o2, 2, 32)
			fmt.Println("o2Decimal:", o2Decimal)
		}
	}

	var z []string = bin
	for bit := 0; bit < 12; bit++ {
		//get the bitcriteria
		_, bitcriteria_epsilon := compute_shit(z)
		fmt.Println(bit, "<-- INDEX THE BIT WE'RE CHECKING FOR CRITERIA")
		fmt.Println("Bitcriteria e: \t", bitcriteria_epsilon)
		for a := 0; a < len(z)-1; a++ {
			fmt.Println("Chunk we're on: \t", z[a])

			if z[a][bit] != bitcriteria_epsilon[bit] {
				fmt.Println("------- SHRINKING THE ARRAY ---------")
				s = append(z[:a], z[a+1:]...)
			}

			o2 = z[a]
			if bit == 12 && a < 2 {
				fmt.Println(z)
			}
			epsDecimal, _ := strconv.ParseInt(o2, 2, 32)
			fmt.Println("epsDecimal:", epsDecimal)
		}

		//		fmt.Println(epsDecimal * o2Decimal)

	}
}

/*
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
*/

func input_handler(filename string) []string {
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
			gamma += "1"
			epsilon += "0"
		}
		zeros = 0
		ones = 0
	}
	return gamma, epsilon
}
