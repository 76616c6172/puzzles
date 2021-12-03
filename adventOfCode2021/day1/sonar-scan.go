package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// read in newline separated numbers, spit out as slice of integers
func input_from_file(filename string) []int {

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	content := string(f)
	d := strings.Split(content, "\n")

	var result []int
	var number int
	for _, s := range d {
		number, _ = strconv.Atoi(s)
		result = append(result, number)
	}

	return result
}

func main() {
	d := input_from_file("data")
	length := len(d)

	// compute the result
	count := 0
	for i := 1; i < length; i++ {
		if d[i] > d[i-1] {
			count++
		}

	}
	fmt.Println(count)
}
