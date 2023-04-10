package main

import "fmt"

func main() {
	// sum elements of array recursively
	fmt.Println(recursiveSum([]int{2, 4, 6, 500}))

	// count elements of array recursively
	fmt.Println(recursiveCountItems([]int{2, 4, 6, 500, 6, 9, 10}))

	// find the max number in array
	fmt.Println(maxNumber([]int{2, 4, 6, 500, 699, 9, 10}))
}

func recursiveSum(s []int) int {
	// base case
	if len(s) == 0 {
		return 0
	}
	// base case
	if len(s) == 1 {
		return s[0]
	}
	// recursive case
	return s[0] + recursiveSum(s[1:])
}

func recursiveCountItems(s []int) int {
	// base case
	if len(s) == 0 {
		return 0
	}
	// base case
	if len(s) == 1 {
		return 1
	}
	// recursive case
	return 1 + recursiveCountItems(s[1:])
}

func maxNumber(s []int) int {
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}
