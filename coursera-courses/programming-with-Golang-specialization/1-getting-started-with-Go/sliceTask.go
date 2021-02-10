package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var intSlice []int
	var input string
	fmt.Println("for exit press enter \"X\"")
	fmt.Scanln(&input)

	for input != "X" {
		intInput, error := strconv.Atoi(input)
		if error != nil {
			fmt.Printf("Error reading input from console: %s", error)
		}
		intSlice = append(intSlice, intInput)
		sort.Ints(intSlice)
		fmt.Println(intSlice)
		fmt.Scanln(&input)
	}
}