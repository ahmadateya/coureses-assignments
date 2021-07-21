package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please input a list of integers delimited by new lines (ENTER between each). When you are done enter X as input")

	input := make([]int, 0)
	var str string
	for {
		fmt.Scan(&str)
		if strings.ToLower(str) == "x" {
			break
		}

		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Please only enter integers")
			return
		}

		input = append(input, num)
	}

	// Get approximate partition size
	partSize := len(input) / 4
	mod := len(input) % 4

	c := make(chan []int)

	startIndex := 0
	for i := 0; i < 4; i++ {
		incrementEnd := 0
		if mod != 0 {
			incrementEnd += 1
			mod -= 1
		}

		// Start goroutine to sort sub array
		go subSort(input[startIndex:startIndex+partSize+incrementEnd], c)
		startIndex += partSize + incrementEnd
	}

	// Listen for channel input for each of the subroutines. This is an unbuffered channel to ensure synchronization
	part1 := <-c
	part2 := <-c
	part3 := <-c
	part4 := <-c

	// Merge arrays into one
	mergedArr := mergeArrays(part1, part2, part3, part4)

	// Sort large array
	sort.Ints(mergedArr)

	// Print final array
	fmt.Printf("Final sorted array: %v\n", mergedArr)
}

func subSort(input []int, c chan []int) {
	fmt.Printf("Sub-array to be sorted: %v\n", input)
	sort.Ints(input)
	c <- input
}

func mergeArrays(arr ...[]int) []int {
	var bigArray []int
	for i := 0; i < len(arr); i++ {
		bigArray = append(bigArray, arr[i]...)
	}

	return bigArray
}