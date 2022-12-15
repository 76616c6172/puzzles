package main

import (
	"fmt"
	"strconv"
	"strings"
)

// create_assignment_pairs expects a string like this:
// 2-4,6-8
// 2-3,4-5
//
// then encodes end returns it as:
// [
//
//	[2, 4] , [6, 8] ]
//	[2, 3] , [4, 5] ]
//
// ]
func create_assignment_pairs(s string) [][2][2]int {
	list := strings.Split(s, "\n")
	var assignmentPairs [][2][2]int
	for _, sectionPair := range list {
		a := strings.Split(sectionPair, ",")
		start1, _ := strconv.Atoi(strings.Split(a[0], "-")[0])
		end1, _ := strconv.Atoi(strings.Split(a[0], "-")[1])
		start2, _ := strconv.Atoi(strings.Split(a[1], "-")[0])
		end2, _ := strconv.Atoi(strings.Split(a[1], "-")[1])
		assignmentPair := [2][2]int{
			{start1, end1},
			{start2, end2},
		}
		assignmentPairs = append(assignmentPairs, assignmentPair)
	}
	return assignmentPairs
}

func doesFullyContainRange(startA, endA, startB, endB int) bool {
	return (startA <= startB && endB <= endA) || (startB <= startA && endA <= endB)
}

// doesExistInRage checks if the needle is contained in the haystack
func doesExistInRange(needle int, haystack []int) bool {
	for _, element := range haystack {
		if element == needle {
			return true
		}
	}
	return false
}

// doesPartiallyOverlap checks if any numbers in the intervallA are also in intervallB
// rangeA := []int{2, 4}
// rangeB := []int{3, 4}
func doesPartiallyOverlap(startA, endA, startB, endB int) bool {
	var rangeA []int
	for i := startA; i <= endA; i++ {
		rangeA = append(rangeA, i)
	}
	var rangeB []int
	for i := startB; i <= endB; i++ {
		rangeB = append(rangeB, i)
	}
	for _, num := range rangeA {
		if doesExistInRange(num, rangeB) {
			return true
		}
	}
	return false
}

func solveDay4Part1(s string) int {
	data := create_assignment_pairs(s)
	var count int
	for i := range data {
		if doesFullyContainRange(data[i][0][0], data[i][0][1], data[i][1][0], data[i][1][1]) {
			count++
		}
	}
	return count
}

func solveDay4Part2(s string) int {
	data := create_assignment_pairs(s)
	var count int
	for i := range data {
		if doesPartiallyOverlap(data[i][0][0], data[i][0][1], data[i][1][0], data[i][1][1]) {
			count++
		}
	}
	return count
}

func solve_day4() {
	filepath := "puzzle-inputs/day4-input"
	input := get_input(filepath)
	solution1 := solveDay4Part1(input)
	solution2 := solveDay4Part2(input)
	fmt.Println("Day 4: Camp Cleanup =", solution1)
	fmt.Println("Day 4: Part Two =", solution2)
}
