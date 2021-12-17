package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Holds the data for each line segment
type line struct {
	start [2]int //line.start[x1][y1]
	end   [2]int //line.end[x2][y2]
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

func data_wrangler(input string) {
	line_amount := strings.Count(input, "\n") + 1
	lines := make([]line, line_amount)
	each_segment := strings.Split(input, "\n")

	// TODO: instead of just printing, store each point coordinate pair correctly in lines[] for each line
	for i := 0; i < line_amount; i++ {
		b := strings.Split(each_segment[i], " -> ")
		fmt.Println(b)
		for z, v := range b {
			point := strings.Split(v, ",")
			if z == 0 {
				fmt.Println("START:", point[0], point[1])
			} else {
				fmt.Println("END:", point[0], point[1])

			}
		}
	}

	fmt.Println(lines)
}

func main() {

	filename := "example"
	test := get_input(filename)
	data_wrangler(test)

}
