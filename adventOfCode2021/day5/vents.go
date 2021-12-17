package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type line struct {
	start [2]int
	end   [2]int
}

// Reads file and returns the contents as a single string with the last newline removed
func get_input(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	result := strings.TrimSuffix(string(f), "\n")
	return string(result)
}

func main() {

	filename := "data"
	test := get_input(filename)

	fmt.Println(test)

}
