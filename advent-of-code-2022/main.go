package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("running main..")
	solve_day1()
}

func getPuzzleInputFromFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(file)
}
