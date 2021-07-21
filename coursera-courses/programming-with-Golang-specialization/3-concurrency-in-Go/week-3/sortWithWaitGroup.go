package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func takeInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter space separated list of integers you want to sort:")

	scanner.Scan()

	return strings.Fields(scanner.Text())
}

func subSort(wg *sync.WaitGroup, nums []int) {
	sort.Ints(nums)
	if wg != nil {
		wg.Done()
	}
}

func main() {
	nums := []int{}

	for _, inputNum := range takeInput() {
		num, _ := strconv.Atoi(inputNum)
		nums = append(nums, num)
	}

	if len(nums) < 4 {
		subSort(nil, nums)
		fmt.Println("Sorted array is:", nums)
		return
	}

	// could be enhanced
	chunkSize := len(nums) / 4

	chunk1 := nums[ : chunkSize] 

	chunk2 := nums[chunkSize : (2 * chunkSize)] 

	chunk3 := nums[(2 * chunkSize) : (3 * chunkSize)]

	chunk4 := nums[(3 * chunkSize) : ]

	var wg sync.WaitGroup
	wg.Add(4)
	go subSort(&wg, chunk1)
	go subSort(&wg, chunk2)
	go subSort(&wg, chunk3)
	go subSort(&wg, chunk4)
	wg.Wait()

	result := mergeSlices(chunk1, chunk2 ,chunk3 ,chunk4)
	// Sort the result
	sort.Ints(result)

	fmt.Println("Sorted array is:", result)
}

// un used but worth to look at
func merge(summedSlice, firstSlice, secondSlice []int) {
	i, j, k := 0, 0, 0

	for i < len(firstSlice) && j < len(secondSlice) {
		if firstSlice[i] < secondSlice[j] {
			summedSlice[k] = firstSlice[i]
			i++
		} else {
			summedSlice[k] = secondSlice[j]
			j++
		}
		k++
	}

	for i < len(firstSlice) {
		summedSlice[k] = firstSlice[i]
		i++
		k++
	}

	for j < len(secondSlice) {
		summedSlice[k] = secondSlice[j]
		j++
		k++
	}
}

func mergeSlices(arr ...[]int) []int {
	var bigArray []int
	for i := 0; i < len(arr); i++ {
		bigArray = append(bigArray, arr[i]...)
	}
	return bigArray
}