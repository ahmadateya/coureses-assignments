package main

import "fmt"

func main() {
	// []int{1, 2, 3, 4, 5, 6, 10, 14, 200, 300}
	sorted := selectionSort([]int{1, 2, 8, 4, 60, 6, 10, 14, 200, 300})
	fmt.Println(sorted)
	fmt.Printf("sorted array: %+v\n", sorted)
}

func selectionSort(arr []int) []int {
	smallestIndex := 0
	sortedArr := []int{len(arr)}
	for i, val := range arr {
		fmt.Printf("============ i: %v, v: %v\n", i, val)
		fmt.Printf("============ val: %v, arr[smallestIndex]: %v\n", val, arr[smallestIndex])

		if val < arr[smallestIndex] {
			smallestIndex = i
		}
		// add to the new array
		sortedArr = append(sortedArr, arr[smallestIndex])
		// delete from old array
		arr = remove(arr, arr[smallestIndex])
		fmt.Println("========== ", sortedArr)
	}
	return sortedArr
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
