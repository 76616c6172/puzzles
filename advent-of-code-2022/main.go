package main

import (
	"fmt"
	"os"
)

func get_input(path string) string {
	filecontent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(filecontent)
}

// *** Entry Point ***
func main() {
	fmt.Println("running main..")
	solve_day1()
	solve_day2()
	solve_day3()
	solve_day4()
}
