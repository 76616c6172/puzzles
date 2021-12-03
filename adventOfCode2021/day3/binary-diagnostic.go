package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	/*
		// datatype to hold each chunck of this pretend binary
		// I will make an array of theese so it can be accessed in this way:
		//
		// [n] abcde
		// ---------
		// [0] 00100
		// [1] 11110
		// [2] 10110
		//
		// to access leftmost "bit" of the first "chunck" bin[0].a
		type bin struct {
			a int //1
			b int //2
			c int //3
			d int //4
			e int //5
			f int //6
			g int //7
			h int //8
			i int //9
			j int //10
			k int //11
			l int //12
		}
	*/

	// or maybe just use *a slice..
	// since the raw input data will be a string we might as well just keep it so
	// read in the data
	var bin []string = input("testdata")
	/*
		for _, a := range bin {
			fmt.Println(a)
		}
	*/

	// I am amazed that this works for accessing the 5th "bit" of the first chunk by doing
	// n := bin[0][4]
	// fmt.Printf("%c", bin[0][4])

	// -------------------------------------------------------------
	// build the answer string
	var answer string
	zeros := 0
	ones := 0

	//FIXME: now what, how do we do this and build the answer string?
	var sw int
	for a, _ := range bin {
		for i := 0; i < 12; i++ {

		}
	}
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
