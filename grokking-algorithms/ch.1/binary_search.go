package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 10, 14, 200, 300}

	// recursive binary search
	i := recursiveBinarySearch(list, 300, 0, len(list)-1)
	if i == -1 {
		fmt.Println("Recursively: Not found!")
	} else {
		fmt.Printf("Recursively: Found at index: %v\n", i)
	}
}
func recursiveBinarySearch(list []int, item, low, high int) (index int) {
	mid := (high + low) / 2
	if low > high { // base case
		index = -1
	} else { // recursive case
		if item < list[mid] {
			return recursiveBinarySearch(list, item, low, mid-1)
		} else if item > list[mid] {
			return recursiveBinarySearch(list, item, mid+1, high)
		} else if item == list[mid] {
			index = mid
		} else {
			index = -1
		}
	}
	return
}
